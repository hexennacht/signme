package main

import (
	"database/sql"
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/hexennacht/signme/services/sm-user-api/config"
	"github.com/hexennacht/signme/services/sm-user-api/core/services"
	"github.com/hexennacht/signme/services/sm-user-api/ent"
	"github.com/hexennacht/signme/services/sm-user-api/ent/migrate"
	"github.com/hexennacht/signme/services/sm-user-api/repository/credential"
	"github.com/hexennacht/signme/services/sm-user-api/repository/user"
	"github.com/hexennacht/signme/services/sm-user-api/server"
	"github.com/hexennacht/signme/services/sm-user-api/server/handler"
	"github.com/hexennacht/signme/services/sm-user-api/utils/ctxutil"
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

// Open new connection
func openPostgresSQL(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func newHandler(conf *config.Configuration) *server.PropertyServer {
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName)

	databaseConnection := openPostgresSQL(databaseURL)

	ctx := ctxutil.NewContextWithConfig(conf)

	schema := databaseConnection.Schema
	if err := schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		panic(err)
	}

	property := new(server.PropertyServer)

	credRepo := credential.NewRepository(databaseConnection.Credential)
	userRepo := user.NewRepository(databaseConnection.User)

	authService := services.NewAuthService(userRepo, credRepo)

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
