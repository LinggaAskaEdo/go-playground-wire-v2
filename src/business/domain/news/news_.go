package news

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func (n *NewsDomain) FindNewsByID(ctx context.Context, newsID int64) (entity.News, error) {
	entityNews := entity.News{}

	tx, err := n.mysql.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return entityNews, common.WrapWithCode(err, http.StatusInternalServerError, "FindNewsByID")
	}

	tx, entityNews, err = n.getSQLNewsByID(ctx, tx, newsID)
	if err != nil {
		tx.Rollback()

		return entityNews, err
	}

	if err := tx.Commit(); err != nil {
		return entityNews, common.WrapWithCode(err, http.StatusInternalServerError, "FindNewsByID")
	}

	return entityNews, nil
}
