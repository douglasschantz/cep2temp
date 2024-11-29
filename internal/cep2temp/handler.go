package cep2temp

import (
	"net/http"

	"github.com/douglasschantz/cep2temp/internal/platform/errors"
	"github.com/gorilla/mux"
)

type Handler interface {
	GetTemperatureByCEP(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func (h handler) GetTemperatureByCEP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	cep := vars["cep"]

	temp, err := h.service.GetTemperatureByCEP(r.Context(), NewRequest(cep))
	if err != nil {
		appErr := errors.Encode(err)
		w.WriteHeader(appErr.Code)
		w.Write(appErr.ToJSON())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(temp.ToJSON())
	return
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
