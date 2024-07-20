package grpc

import (
	"context"

	"github.com/jcserv/mjurl/internal/transport/grpc/pb"
	"github.com/jcserv/mjurl/model"
)

type GRPC struct {
}

func NewGRPC(dependencies model.Dependencies) *GRPC {
	return &GRPC{}
}

func (g *GRPC) GetURL(ctx context.Context, shortURL *pb.ShortURL) (*pb.LongURL, error) {
	return nil, nil
}

func (g *GRPC) ShortenURL(ctx context.Context, longURL *pb.LongURL) (*pb.ShortURL, error) {
	return nil, nil
}
