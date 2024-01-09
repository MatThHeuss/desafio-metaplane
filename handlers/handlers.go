package handlers

import (
	"encoding/json"
	"github.com/MatThHeuss/desafio-metaplane/listOperations"
	"net/http"
)

var lists []listoperations.List // Armazenar listas

// Handler para salvar listas
func SaveListsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newList listoperations.List
	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if len(lists) >= 2 {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Maximum size reached"})
		return
	}

	lists = append(lists, newList)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "List saved successfully"})
}

// Handler para mesclar listas
func MergeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(lists) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Need at least two lists to merge"})
		return
	}

	l1 := listoperations.SliceToListNode(lists[0].Numbers)
	l2 := listoperations.SliceToListNode(lists[1].Numbers)
	mergedList := listoperations.MergeTwoLists(l1, l2)

	var mergedNumbers []int
	for mergedList != nil {
		mergedNumbers = append(mergedNumbers, mergedList.Val)
		mergedList = mergedList.Next
	}
	lists = lists[:0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mergedNumbers)
}
