package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/nabilaps1/API-POSTGRESQL/models"
)

// chi é um pacote q ajuda a capturar valores (parametros) das rotas

func Delete(w http.ResponseWriter, r *http.Request) { // recebe um id (na url) como parametro
	// todos os valores rescebidos pela URL sao string. é preciso converter id para int
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// se nao há erros, atualiza os dados no banco de dados
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// validacao da atualizacao para apenas 1 registro
	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros", rows)

	}

	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso!",
	}

	// resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp) // faz o encode do map para um json de resposta da aplicacao
}
