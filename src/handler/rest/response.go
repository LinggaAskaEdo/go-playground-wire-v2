package rest

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/entity"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

const (
	timeFormat = "Monday, 02 January 2006 15:04 MST"
)

type NewsResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Content       string `json:"content"`
	Summary       string `json:"summary"`
	ArticleTS     int64  `json:"article_ts"`
	PublishedDate string `json:"published_date"`
	Inserted      string `json:"created_at"`
	Updated       string `json:"updated_at"`
}

func (e *Rest) httpRespError(w http.ResponseWriter, r *http.Request, response *common.WebResponse) {
	response.Path = r.URL.String()

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	common.FatalIfError(err)
}

func (e *Rest) httpRespSuccess(w http.ResponseWriter, r *http.Request, response interface{}) {
	bodyResp := common.WebResponse{}

	switch data := response.(type) {
	case entity.News:
		bodyResp = common.WebResponse{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Path:   r.URL.String(),
			Data:   e.parseNewsReponse(&data),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(bodyResp)
	common.FatalIfError(err)
}

func (e *Rest) parseNewsReponse(data *entity.News) NewsResponse {
	return NewsResponse{
		ID:            data.ID,
		Title:         data.Title,
		URL:           data.URL,
		Content:       data.Content,
		Summary:       e.isNullString(data.Summary),
		ArticleTS:     data.ArticleTS,
		PublishedDate: e.isNullTime(data.PublishedDate),
		Inserted:      e.isNullTime(data.Inserted),
		Updated:       e.isNullTime(data.Updated),
	}
}

func (e *Rest) isNullString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}

	return s.String
}

func (e *Rest) isNullTime(s sql.NullTime) string {
	if !s.Valid {
		return ""
	}

	return s.Time.Format(timeFormat)
}
