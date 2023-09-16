package ctxutil

import (
	"context"

	"github.com/hexennacht/signme/services/sm-user-api/config"
)

func NewContextWithConfig(conf *config.Configuration) context.Context {
	return context.WithoutCancel(context.WithValue(context.Background(), "config", conf))
}
