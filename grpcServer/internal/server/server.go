package server

import (
	"context"

	api "github.com/wainola/proglog/grpcServer/api/v1"
)

var _ api.LogServer = (*grpcServer)(nil)

type Config struct {
	ComitLog ComitLog
}

type grpcServer struct {
	*Config
}

func newgrpcServer(config *Config) (srv *grpcServer, err error) {
	srv = &grpcServer{
		Config: config,
	}

	return srv, nil
}

func (s *grpcServer) Produce(ctx context.Context, req *api.ProduceRequest) (*api.ProduceResponse, error) {
	offset, err := s.Config.Append(req.Record)

	if err != nil {
		return nil, err
	}

	return &api.ProduceResponse{Offset: offset}, nil
}

func (s *grpcServer) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	record, err := s.ComitLog.Read(req.Offset)

	if err != nil {
		return nil, err
	}

	return &api.ConsumeResponse{Record: record}, nil
}
