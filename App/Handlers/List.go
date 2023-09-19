package Handlers

import (
	"encoding/json"
	"golang-start/App/App/Models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := Models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)

	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}