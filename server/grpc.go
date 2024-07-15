package server

import (
	"net"

	"google.golang.org/grpc"
)

var _ Server = (*GRPCServer)(nil)

// Server grpc server
type GRPCServer struct {
	srv  *grpc.Server
	addr string
}

// NewServer creates a new server instance with default settings
func NewGRPCServer(grpcServer *grpc.Server, config *Config) *GRPCServer {
	ser := GRPCServer{
		srv:  grpcServer,
		addr: config.Addr(),
	}

	return &ser
}

// Start to start the server and listen on the given address
func (h *GRPCServer) Start() (err error) {
	lis, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}
	if err = h.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

// Shutdown shuts down the server
func (h *GRPCServer) Shutdown() error {
	h.srv.GracefulStop()
	return nil
}
