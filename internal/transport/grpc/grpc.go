package grpc

import (
	v1 "github.com/jcserv/mjurl/internal/transport/grpc/v1"
	"github.com/jcserv/mjurl/model"
	"google.golang.org/grpc"
)

type GRPC struct {
	V1Server *v1.Server
}

func NewGRPC(dependencies model.Dependencies) *GRPC {
	return &GRPC{
		V1Server: v1.NewServer(dependencies.URLService),
	}
}

func (g *GRPC) RegisterServer(server *grpc.Server) {
	g.V1Server.RegisterServer(server)
}
