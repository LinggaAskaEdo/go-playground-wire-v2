package usecase

import (
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/domain"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/usecase/news"
)

type Usecase struct {
	News news.UsecaseItf
}

type Options struct {
	News news.Options
}

func Init(
	options Options,
	dom *domain.Domain,
) *Usecase {
	return &Usecase{
		News: news.InitNewsUsecase(
			options.News,
			dom.News,
		),
	}
}
