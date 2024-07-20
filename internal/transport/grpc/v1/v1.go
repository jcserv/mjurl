package v1

import (
	"context"

	"github.com/jcserv/mjurl/internal/transport/grpc/pb"
	"github.com/jcserv/mjurl/internal/url"
	"github.com/jcserv/mjurl/model"
	"google.golang.org/grpc"
)

type Server struct {
	URLService model.IURLService
	pb.UnimplementedMJUrlServer
}

func NewServer(urlService model.IURLService) *Server {
	return &Server{URLService: urlService}
}

func (s *Server) RegisterServer(server *grpc.Server) {
	pb.RegisterMJUrlServer(server, s)
}

func (s *Server) ShortenURL(ctx context.Context, longURL *pb.LongURL) (*pb.ShortURL, error) {
	command, err := url.NewShortenURL(*longURL.LongURL)
	if err != nil {
		return nil, nil
	}

	url, err := command.Execute(ctx, s.URLService)
	if err != nil {
		return nil, nil
	}
	return &pb.ShortURL{
		ShortURL: (*string)(&url.Short),
	}, nil
}

func (s *Server) GetURL(ctx context.Context, shortURL *pb.ShortURL) (*pb.LongURL, error) {
	command, err := url.NewGetURLByShort(*shortURL.ShortURL)
	if err != nil {
		return nil, err
	}

	url, err := command.Execute(ctx, s.URLService)
	if err != nil {
		return nil, err
	}

	return &pb.LongURL{
		LongURL: (*string)(&url.Long),
	}, nil
}
