package news

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/domain/news"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
)

type UsecaseItf interface {
	FindNewsByID(ctx context.Context, newsID int64) (entity.News, error)
}

type NewsUsecase struct {
	opt  Options
	news news.DomainItf
}

type Options struct {
	RssUrl string
}

func InitNewsUsecase(
	opt Options,
	news news.DomainItf,
) UsecaseItf {
	return &NewsUsecase{
		opt:  opt,
		news: news,
	}
}
