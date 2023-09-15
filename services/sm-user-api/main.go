package main

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/hexennacht/signme/services/sm-user-api/config"
	"github.com/hexennacht/signme/services/sm-user-api/core/services"
	"github.com/hexennacht/signme/services/sm-user-api/server"
	"github.com/hexennacht/signme/services/sm-user-api/server/handler"
)

var (
	id, _ = os.Hostname()
)

func main() {
	conf := config.ReadConfig()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", conf.AppName,
		"service.version", conf.AppVersion,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	property := newHandler(conf)

	grpcServer := server.NewGRPCServer(conf, property)
	httpServer := server.NewHTTPServer(conf, property)

	app := newApp(logger, conf, grpcServer, httpServer)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newHandler(conf *config.Configuration) *server.PropertyServer {
	property := new(server.PropertyServer)

	authService := services.NewAuthService()

	property.AuthHandler = handler.NewAuthHandler(authService)

	return property
}

func newApp(logger log.Logger, conf *config.Configuration, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(conf.AppName),
		kratos.Version(conf.AppVersion),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}
