package handlers

import (
	"bytes"
	"encoding/json"
	listoperations "github.com/MatThHeuss/desafio-metaplane/listOperations"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	setup := func() {
		lists = nil // Redefinir a variável global 'lists'
	}
	t.Run("Should save a list", func(t *testing.T) {
		setup()
		list := listoperations.List{Numbers: []int{1, 2, 3}}
		body, _ := json.Marshal(list)
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Len(t, lists, 1)
		assert.Equal(t, list, lists[0])
	})

	t.Run("SHould return an error when trying to save more than 2 lists", func(t *testing.T) {
		// Limpar estado
		setup()

		// Preparar estado inicial
		lists = append(lists, listoperations.List{Numbers: []int{1, 2, 3}})
		lists = append(lists, listoperations.List{Numbers: []int{4, 5, 6}})

		// Tenta adicionar uma terceira lista
		newList := listoperations.List{Numbers: []int{7, 8, 9}}
		body, _ := json.Marshal(newList)
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		// Executa o handler
		handler.ServeHTTP(rr, req)

		// Verificar se o código de status correto foi retornado
		assert.Equal(t, http.StatusForbidden, rr.Code, "Should return 403 Forbidden for maximum size reached")

		// Verificar se a mensagem de erro corresponde
		var resp map[string]string
		err := json.NewDecoder(rr.Body).Decode(&resp)
		if err != nil {
			t.Fatalf("Erro ao decodificar a resposta: %v", err)
		}
		assert.Equal(t, "Maximum size reached", resp["error"], "Should return correct error message")
	})

	t.Run("Should return an error when trying to encode invalid json", func(t *testing.T) {
		invalidJSON := []byte("{invalid-json}")
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(invalidJSON))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		// Executa o handler
		handler.ServeHTTP(rr, req)

		// Verifica se o status BadRequest foi retornado
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Should return 400 BadRequest for invalid JSON")

		// Verifica se a mensagem de erro é a esperada
		var resp map[string]string
		err := json.NewDecoder(rr.Body).Decode(&resp)
		if err != nil {
			t.Fatalf("Erro ao decodificar a resposta: %v", err)
		}
		assert.Contains(t, resp["error"], "invalid character", "Should contain error message about invalid JSON")
	})

	t.Run("Should merge two lists", func(t *testing.T) {
		setup()
		list01 := listoperations.List{Numbers: []int{1, 3, 4, 7}}
		list02 := listoperations.List{Numbers: []int{2, 5, 6}}
		resultList := []int{1, 2, 3, 4, 5, 6, 7}

		lists = append(lists, list01)
		lists = append(lists, list02)
		req, _ := http.NewRequest("POST", "/merge", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(MergeHandler)

		handler.ServeHTTP(rr, req)
		var gotNumbers []int
		err := json.NewDecoder(rr.Body).Decode(&gotNumbers)
		if err != nil {
			t.Fatal("Failed to decode response body")
		}

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, resultList, gotNumbers)
	})

	t.Run("SHould return error when lists length < 2", func(t *testing.T) {
		setup() // Limpa o estado, definido em outro comentário
		req, _ := http.NewRequest("POST", "/merge", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(MergeHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "Should return 400 Bad Request for not enough lists")
		assert.Contains(t, rr.Body.String(), "Need at least two lists to merge", "Should contain error message")
	})

}
