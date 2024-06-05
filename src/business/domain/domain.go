package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/domain/news"
)

type Domain struct {
	News news.DomainItf
}

type Options struct {
	News news.Options
}

func Init(
	option Options,
	mysql *sql.DB,
	redis *redis.Client,
) *Domain {
	return &Domain{
		News: news.InitNewsDomain(
			option.News,
			mysql,
			redis,
		),
	}
}
