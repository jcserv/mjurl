package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jcserv/mjurl/internal/transport/api"
	mjurlGRPC "github.com/jcserv/mjurl/internal/transport/grpc"
	"github.com/jcserv/mjurl/internal/url"
	"github.com/jcserv/mjurl/internal/utils/log"
	"github.com/jcserv/mjurl/model"
	"google.golang.org/grpc"
)

type MJURLService struct {
	api  *api.API
	grpc *mjurlGRPC.GRPC
	cfg  *Configuration
}

func NewMJURLService() (*MJURLService, error) {
	cfg, err := NewConfiguration()
	if err != nil {
		return nil, err
	}

	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	urlService := url.NewURLService(url.NewPSQLStore(dbpool))
	deps := model.Dependencies{URLService: urlService}
	api := api.NewAPI(deps)
	grpc := mjurlGRPC.NewGRPC(deps)

	s := &MJURLService{
		api,
		grpc,
		cfg,
	}

	return s, nil
}

// Run starts the service
func (s *MJURLService) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		s.StartHTTP(ctx)
	}(ctx)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		s.StartGRPC(ctx)
	}(ctx)

	wg.Wait()
	return nil
}

// Shutdown shuts down the service
func (s *MJURLService) Shutdown() {
}

func (s *MJURLService) StartHTTP(ctx context.Context) error {
	log.Info(ctx, fmt.Sprintf("Starting HTTP server on port %s", s.cfg.HTTPPort))
	r := s.api.RegisterRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.HTTPPort), r)
	return nil
}

func (s *MJURLService) StartGRPC(ctx context.Context) error {
	log.Info(ctx, fmt.Sprintf("Starting GRPC server on port %s", s.cfg.GRPCPort))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.cfg.GRPCPort))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	s.grpc.RegisterServer(server)

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
