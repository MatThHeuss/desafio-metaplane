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
	t.Run("Should successfully save a new list via POST request and return 201 Created", func(t *testing.T) {
		// Arrange
		setup()
		list := listoperations.List{Numbers: []int{1, 2, 3}}
		body, _ := json.Marshal(list)
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Len(t, lists, 1)
		assert.Equal(t, list, lists[0])
	})

	t.Run("Should return 403 Forbidden when trying to save a third list, exceeding the limit of 2", func(t *testing.T) {
		// Arrange
		setup()
		lists = append(lists, listoperations.List{Numbers: []int{1, 2, 3}})
		lists = append(lists, listoperations.List{Numbers: []int{4, 5, 6}})

		newList := listoperations.List{Numbers: []int{7, 8, 9}}
		body, _ := json.Marshal(newList)
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusForbidden, rr.Code, "Should return 403 Forbidden for maximum size reached")
		assert.Contains(t, rr.Body.String(), "Maximum size reached", "Should return correct error message")
	})

	t.Run("Should return 400 BadRequest when a POST request with invalid JSON is received", func(t *testing.T) {
		// Arrange
		invalidJSON := []byte("{invalid-json}")
		req, _ := http.NewRequest("POST", "/savelists", bytes.NewBuffer(invalidJSON))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SaveListsHandler)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Should return 400 BadRequest for invalid JSON")
		assert.Contains(t, rr.Body.String(), "invalid character", "Should contain error message about invalid JSON")
	})

	t.Run("Should merge two lists into a single sorted list and return 200 OK", func(t *testing.T) {
		// Arrange
		setup()
		list01 := listoperations.List{Numbers: []int{1, 3, 4, 7}}
		list02 := listoperations.List{Numbers: []int{2, 5, 6}}
		resultList := []int{1, 2, 3, 4, 5, 6, 7}

		lists = append(lists, list01)
		lists = append(lists, list02)
		req, _ := http.NewRequest("POST", "/merge", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(MergeHandler)

		// Act
		handler.ServeHTTP(rr, req)
		var gotNumbers []int
		err := json.NewDecoder(rr.Body).Decode(&gotNumbers)
		if err != nil {
			t.Fatal("Failed to decode response body")
		}

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, resultList, gotNumbers)
	})

	t.Run("Should return 400 Bad Request when attempting to merge with fewer than two lists", func(t *testing.T) {
		// Arrange
		setup()
		req, _ := http.NewRequest("POST", "/merge", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(MergeHandler)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, rr.Code, "Should return 400 Bad Request for not enough lists")
		assert.Contains(t, rr.Body.String(), "Need at least two lists to merge", "Should contain error message")
	})

	t.Run("", func(t *testing.T) {
		// Arrange
		req, _ := http.NewRequest("GET", "/health", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthCheckHandler)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code, "Status code should be 200")

		contentType := rr.Header().Get("Content-Type")
		assert.Equal(t, "application/json", contentType, "Content type should be application/json")

		assert.Contains(t, rr.Body.String(), "ok")
	})

}
