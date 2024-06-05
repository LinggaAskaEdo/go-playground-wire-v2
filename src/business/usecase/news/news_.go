package news

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
)

func (n *NewsUsecase) FindNewsByID(ctx context.Context, newsID int64) (entity.News, error) {
	news, err := n.news.FindNewsByID(ctx, newsID)
	if err != nil {
		return news, err
	}

	return news, nil
}
