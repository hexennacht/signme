package server

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/hexennacht/signme/services/sm-user-api/config"
	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *config.Configuration, property *PropertyServer) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.AppGRPCHost != "" {
		opts = append(opts, grpc.Network(c.AppGRPCHost))
	}
	if c.AppGRPCHost != "" && c.AppGRPCPort > 0 {
		opts = append(opts, grpc.Address(fmt.Sprintf("%s:%d", c.AppGRPCHost, c.AppGRPCPort)))
	}
	if c.AppGRPCTimeout > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.AppGRPCTimeout)))
	}

	srv := grpc.NewServer(opts...)

	auth.RegisterAuthenticationServer(srv, property.AuthHandler)

	return srv
}
