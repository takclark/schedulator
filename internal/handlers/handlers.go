package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/takclark/schedulator/api"
	"github.com/takclark/schedulator/internal/service"
)

type handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *handler {
	return &handler{s: s}
}

func (h *handler) Register() *mux.Router {
	r := mux.NewRouter()
	hr := r.PathPrefix("/hack").Subrouter()

	hr.HandleFunc("/rules/{id}", h.rule).Methods(http.MethodGet)
	hr.HandleFunc("/rules", h.rules).Methods(http.MethodGet)
	hr.HandleFunc("/rules", h.createRule).Methods(http.MethodPost)
	// hr.HandleFunc("/rule/{id}", h.updateRule).Methods(http.MethodPatch)

	return r
}

func (h *handler) rule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rule, err := h.s.Rule(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := respondWithJSON(w, rule); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *handler) rules(w http.ResponseWriter, r *http.Request) {
	rules, err := h.s.Rules()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := respondWithJSON(w, rules); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *handler) createRule(w http.ResponseWriter, r *http.Request) {
	data := api.CreateRule{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rule, err := h.s.CreateRule(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := respondWithJSON(w, rule); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func respondWithJSON(w http.ResponseWriter, t any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(t); err != nil {
		return err
	}

	return nil
}
