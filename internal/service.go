package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jcserv/mjurl/internal/transport/api"
	"github.com/jcserv/mjurl/internal/url"
	"github.com/jcserv/mjurl/internal/utils/log"
)

type MJURLService struct {
	api *api.API
	cfg *Configuration
}

func NewMJURLService() (*MJURLService, error) {
	cfg, err := NewConfiguration()
	if err != nil {
		return nil, err
	}

	urlService := url.NewURLService()
	api := api.NewAPI(urlService)

	s := &MJURLService{
		api,
		cfg,
	}

	return s, nil
}

// Run starts the service
func (s *MJURLService) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return s.StartHTTP(ctx)
}

// Shutdown shuts down the service
func (s *MJURLService) Shutdown() {
}

func (s *MJURLService) StartHTTP(ctx context.Context) error {
	log.Info(ctx, fmt.Sprintf("Starting HTTP server on port %s", s.cfg.Port))
	r := s.api.RegisterRoutes()
	http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.Port), r)
	return nil
}
