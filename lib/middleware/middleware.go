package middleware

import (
	"context"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

type Middleware struct {
	Handler http.Handler
}

func InitMiddleware(handler http.Handler) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}

func (middleware *Middleware) ServeHTTP(ctx context.Context, writer http.ResponseWriter, request *http.Request) {
	config := ctx.Value("token").(common.TokenConfiguration)

	if config.Auth == request.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := common.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		common.WriteToResponseBody(writer, webResponse)
	}
}
