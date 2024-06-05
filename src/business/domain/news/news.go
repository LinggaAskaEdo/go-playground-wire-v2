package news

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
	"github.com/redis/go-redis/v9"
)

type DomainItf interface {
	FindNewsByID(ctx context.Context, categoryId int64) (entity.News, error)
}

type NewsDomain struct {
	opt   Options
	mysql *sql.DB
	redis *redis.Client
}

type Options struct {
}

func InitNewsDomain(
	opt Options,
	mysql *sql.DB,
	redis *redis.Client,
) DomainItf {
	return &NewsDomain{
		opt:   opt,
		mysql: mysql,
		redis: redis,
	}
}
