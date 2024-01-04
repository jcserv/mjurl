package internal

import (
	"context"

	"github.com/jcserv/mjurl/internal/transport/http"
	"github.com/jcserv/mjurl/internal/utils/log"
)

type MJURLService struct {
}

func NewMJURLService() (*MJURLService, error) {
	return &MJURLService{}, nil
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
	http.RegisterRoutes()
	return nil
}
