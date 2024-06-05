package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v2/lib/middleware"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

type ServerOptions struct {
	Middleware *middleware.Middleware
}

func InitServer(ctx context.Context, opts *ServerOptions) *http.Server {
	config := ctx.Value("server").(common.ServerConfiguration)

	address := fmt.Sprintf(
		"%s:%s",
		config.ServerHost,
		config.ServerPort,
	)

	return &http.Server{
		Addr:    address,
		Handler: opts.Middleware.Handler,
	}
}
