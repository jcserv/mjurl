package internal

import (
	"context"
	"net/http"

	"github.com/jcserv/mjurl/internal/transport/api"
	"github.com/jcserv/mjurl/internal/utils/log"
)

type MJURLService struct {
	api *api.API
}

func NewMJURLService() (*MJURLService, error) {
	s := &MJURLService{}
	api := api.NewAPI()
	s.api = api

	return s, nil
}

// Run starts the service
func (s *MJURLService) Run() error {
	ctx, _ := context.WithCancel(context.Background())
	return s.StartHTTP(ctx)
}

// Shutdown shuts down the service
func (s *MJURLService) Shutdown() {
}

func (s *MJURLService) StartHTTP(ctx context.Context) error {
	log.Info(ctx, "HTTP server started on port 8080")
	r := s.api.RegisterRoutes()
	http.ListenAndServe(":8080", r)
	return nil
}
