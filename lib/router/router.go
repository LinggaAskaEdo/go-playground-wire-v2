package router

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	return mux.NewRouter()
}
