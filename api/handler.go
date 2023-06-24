package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gomicroservice/service"
	"net/http"
)

type Handler interface {
	GetRoutes() *mux.Router
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get/{key}", h.GetHandler).Methods("GET")
	r.HandleFunc("/set/{key}/{value}", h.SetHandler).Methods("POST")

	return r
}

func (h *handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := h.service.Get(r.Context(), key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(value)
}

func (h *handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]

	err := h.service.Set(r.Context(), key, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Successfully set value")
}
