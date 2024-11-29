package router

import (
	"net/http"

	"github.com/douglasschantz/cep2temp/internal/cep2temp"
	"github.com/gorilla/mux"
)

func SetupRouter(handler cep2temp.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cep2temp/{cep}", handler.GetTemperatureByCEP).Methods(http.MethodGet)
	return r
}
