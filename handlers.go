package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Register() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/rule/{id}", s.rule)

	return r
}

func (s *Server) rule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	rule, err := s.Rule(id)
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
