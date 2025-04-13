package grpc

import (
	desc "calculator/pkg/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type serverGRPC struct {
	Server *grpc.Server
	lis    net.Listener
}

// NewServerGRPC конструктор для serverGRPC.
func NewServerGRPC() *serverGRPC {
	return &serverGRPC{
		Server: grpc.NewServer(),
	}
}

// Run запускает gRPC сервер.
func (s *serverGRPC) Run(port string) error {
	var err error
	s.lis, err = net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	if err := s.Server.Serve(s.lis); err != nil {
		return err
	}

	return nil
}

// ShutDown останавливает gRPC сервер.
func (s *serverGRPC) ShutDown(ctx context.Context) error {
	stopped := make(chan struct{})
	go func() {
		s.Server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		return nil
	case <-ctx.Done():
		s.Server.Stop()
		return ctx.Err()
	}
}

func Registration(server *grpc.Server, urlService desc.URLServiceServer) {
	desc.RegisterURLServiceServer(server, urlService)
	reflection.Register(server)
}
