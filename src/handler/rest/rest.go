package rest

import (
	"sync"

	"github.com/gorilla/mux"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/business/usecase"
)

var once = &sync.Once{}

type REST interface {
}

type Rest struct {
	mux *mux.Router
	uc  *usecase.Usecase
}

func InitRest(mux *mux.Router, uc *usecase.Usecase) REST {
	var e *Rest

	once.Do(func() {
		e = &Rest{
			mux: mux,
			uc:  uc,
		}

		e.Serve()
	})

	return e
}

func (e *Rest) Serve() {
	e.mux.HandleFunc("/test", e.TestHandler).Methods("GET")
	e.mux.HandleFunc("/api/news/{id}", e.FindNewsByID).Methods("GET")
}
