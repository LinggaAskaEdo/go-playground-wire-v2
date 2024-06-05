package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/domain"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/usecase"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/usecase/news"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/handler/rest"
)

var (
	ctx context.Context

	// scheduler gocron.Scheduler
	httpserver *http.Server
)

func init() {
	ctx = context.Background()

	log.Println("Initializing app config...")
	appConfig, err := InitializeConfig()
	common.PanicIfError(err)

	ctx = context.WithValue(ctx, "mysql", appConfig.Database.MySQL)
	ctx = context.WithValue(ctx, "redis", appConfig.Cache.Redis)
	ctx = context.WithValue(ctx, "token", appConfig.Token)
	ctx = context.WithValue(ctx, "server", appConfig.Server)
	// log.Println(ctx.Value("mysql").(common.MySQLConfiguration).MySQLHost)

	log.Println("Initializing database...")
	db, err := InitializeMySQL(ctx)
	common.PanicIfError(err)

	log.Println("Initializing cache...")
	cache, err := InitializeRedis(ctx)
	common.PanicIfError(err)

	log.Println("Initializing validator...")
	validators, err := InitializeValidator()
	common.PanicIfError(err)

	log.Println("Initializing router...")
	domOpts := domain.Options{}
	ucOpts := usecase.Options{
		News: news.Options{
			RssUrl: appConfig.Business.Usecase.News.RssURL,
		},
	}
	business := InitializeBusiness(domOpts, ucOpts, db, cache)
	httpRouter := InitializeRouter()
	_ = rest.InitRest(httpRouter, business)

	log.Println("Initializing server...")
	httpserver, err = InitializeServer(ctx, db, cache, validators, httpRouter)
	common.PanicIfError(err)

	log.Println("Service started...")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

func main() {
	// scheduler.Start()

	err := httpserver.ListenAndServe()
	common.PanicIfError(err)
}

func cleanup() {
	log.Println("Cleaning up process...")
	// err := scheduler.Shutdown()
	// helper.PanicIfError(err)

	err := httpserver.Shutdown(ctx)
	common.PanicIfError(err)
}
