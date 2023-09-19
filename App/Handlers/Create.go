package Handlers

import (
	"encoding/json"
	"fmt"
	"golang-start/App/App/Models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo Models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := Models.Create(todo)

	var resp map[string]any

	if resp != nil {
		resp = map[string]any {
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any {
			"Error": false,
			"Message": fmt.Sprintf("Todo Inserido Com Sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}