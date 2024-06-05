//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire-v2/lib/cache"
	"github.com/linggaaskaedo/go-playground-wire-v2/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v2/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v2/lib/middleware"
	"github.com/linggaaskaedo/go-playground-wire-v2/lib/router"
	"github.com/linggaaskaedo/go-playground-wire-v2/lib/server"
	valid "github.com/linggaaskaedo/go-playground-wire-v2/lib/validator"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/domain"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/usecase"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func InitializeConfig() (common.Configuration, error) {
	wire.Build(
		config.InitConfig,
	)

	return common.Configuration{}, nil
}

func InitializeMySQL(ctx context.Context) (*sql.DB, error) {
	wire.Build(
		database.InitMySQL,
	)

	return nil, nil
}

func InitializeRedis(ctx context.Context) (*redis.Client, error) {
	wire.Build(
		cache.InitRedis,
	)

	return nil, nil
}

func InitializeValidator() (*validator.Validate, error) {
	wire.Build(
		valid.InitValidator,
	)

	return nil, nil
}

func InitializeBusiness(domOpts domain.Options, ucOpts usecase.Options, mysql *sql.DB, redis *redis.Client) *usecase.Usecase {
	wire.Build(
		domain.Init,
		usecase.Init,
	)

	return nil
}

func InitializeRouter() *mux.Router {
	wire.Build(
		router.InitRouter,
	)

	return nil
}

func InitializeServer(ctx context.Context, db *sql.DB, cache *redis.Client, validator *validator.Validate, router *mux.Router) (*http.Server, error) {
	wire.Build(
		wire.Bind(new(http.Handler), new(*mux.Router)),
		middleware.InitMiddleware,
		wire.NewSet(wire.Struct(new(server.ServerOptions), "*"), server.InitServer),
	)

	return nil, nil
}
