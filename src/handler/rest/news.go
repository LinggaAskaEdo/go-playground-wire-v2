package rest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func (e *Rest) TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorillaxxx !!!\n"))
}

func (e *Rest) FindNewsByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]
	newsID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		e.httpRespError(w, r, common.WrapWithCode(err, http.StatusBadRequest, "Invalid id"))
		return
	}

	news, err := e.uc.News.FindNewsByID(ctx, newsID)
	if err != nil {
		e.httpRespError(w, r, common.WrapWithCode(err, http.StatusNotFound, "Data not found"))
		return
	}

	// webResponse := common.WebResponse{
	// 	Code:   http.StatusOK,
	// 	Status: http.StatusText(http.StatusOK),
	// 	Data:   news,
	// }

	// WriteToResponseBody(w, webResponse)
	e.httpRespSuccess(w, r, news)
}
