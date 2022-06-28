package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/takclark/schedulator/internal/engine"
	"github.com/takclark/schedulator/internal/handlers"
	"github.com/takclark/schedulator/internal/service"
	"github.com/takclark/schedulator/internal/service/database"
)

type Rule struct {
	Expression string `json:"expression,omitempty"`
	Name       string `json:"name,omitempty"`
}

func main() {
	logger := log.Default()

	e := engine.NewEngine(logger)
	if err := e.Start(); err != nil {
		panic(err)
	}

	database, err := database.NewSqlStore("./data/data.db")
	if err != nil {
		panic(err)
	}

	s := service.NewService(e, database, logger)
	h := handlers.NewHandler(s)

	fmt.Println("server starting up...")

	routes := h.Register()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), routes); err != nil {
		panic(err)
	}
}
