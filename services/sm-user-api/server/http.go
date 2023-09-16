package server

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/hexennacht/signme/services/sm-user-api/config"
	"github.com/hexennacht/signme/services/sm-user-api/grpc/v1/auth"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *config.Configuration, property *PropertyServer) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}

	if c.AppHTTPHost != "" {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%d", c.AppHTTPHost, c.AppHTTPPort)))
	}

	if c.AppHTTPTimeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.AppHTTPTimeout)))
	}

	srv := http.NewServer(opts...)

	auth.RegisterAuthenticationHTTPServer(srv, property.AuthHandler)

	return srv
}
