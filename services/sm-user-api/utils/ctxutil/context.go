package ctxutil

import (
	"context"

	"github.com/hexennacht/signme/services/sm-user-api/config"
)

func NewContextWithConfig(conf *config.Configuration) context.Context {
	return context.WithValue(context.Background(), "config", conf)
}
