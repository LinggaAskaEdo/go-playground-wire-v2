package news

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func (n *NewsDomain) getSQLNewsByID(ctx context.Context, tx *sql.Tx, newsID int64) (*sql.Tx, entity.News, error) {
	result := entity.News{}

	rows, err := tx.QueryContext(ctx, GetNewsByID, newsID)
	if err != nil {
		return tx, result, common.WrapWithCode(err, http.StatusInternalServerError, "getSQLNewsByID")
	}

	if rows.Next() {
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&result.URL,
			&result.Content,
			&result.Summary,
			&result.ArticleTS,
			&result.PublishedDate,
			&result.Inserted,
			&result.Updated,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return tx, result, common.WrapWithCode(err, http.StatusInternalServerError, "getSQLNewsByID")
			}

			return tx, result, common.WrapWithCode(err, http.StatusInternalServerError, "getSQLNewsByID")
		}

		return tx, result, nil
	} else {
		return tx, result, common.WrapWithCode(err, http.StatusNotFound, "getSQLNewsByID")
	}
}
