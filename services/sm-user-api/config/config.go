package config

import "github.com/kelseyhightower/envconfig"

type Configuration struct {
	AppName        string `envconfig:"APP_NAME" default:"sm-user-api"`
	AppVersion     string `envconfig:"APP_VERSION" default:"0.0.1"`
	AppHTTPHost    string `envconfig:"APP_HTTP_HOST" default:"127.0.0.1"`
	AppHTTPPort    int    `envconfig:"APP_HTTP_PORT" default:"8000"`
	AppHTTPTimeout int    `envconfig:"APP_HTTP_TIMEOUT" default:"60"`
	AppGRPCHost    string `envconfig:"APP_GRPC_HOST" default:"127.0.0.1"`
	AppGRPCPort    int    `envconfig:"APP_GRPC_PORT" default:"8080"`
	AppGRPCTimeout int    `envconfig:"APP_GRPC_TIMEOUT" default:"60"`

	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"127.0.0.1"`
	DatabasePort     int    `envconfig:"DATABASE_PORT" default:"3306"`
	DatabaseUsername string `envconfig:"DATABASE_USERNAME" default:"root"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"secret"`
	DatabaseName     string `envconfig:"DATABASE_Name" default:"sm_user"`
}

func ReadConfig() *Configuration {
	conf := new(Configuration)

	envconfig.MustProcess("", &conf)

	return conf
}
