package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if payload != nil {
		if erro := json.NewEncoder(w).Encode(payload); erro != nil {
			log.Fatal(erro)
		}
	}

}

func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		err.Error(),
	})

}
