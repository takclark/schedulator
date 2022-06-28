package main

import (
	"fmt"
	"net/http"
)

type Rule struct {
	Expression string `json:"expression,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Server struct {
	Port int
}

func (s *Server) Rule(id string) (Rule, error) {
	r := Rule{
		Expression: "EXP",
		Name:       "Nunny",
	}

	return r, nil
}

func main() {
	s := Server{
		Port: 8080,
	}

	fmt.Println("server starting up...")

	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Register())
}
