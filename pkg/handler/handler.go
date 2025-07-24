package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pluralsh/pr-governance-webhook/api"
)

func Open(in *api.OpenInput) (map[string]any, error) {
	// Example: just echo some state info back
	state := map[string]any{
		"url":   in.Pr.Url,
		"title": in.Pr.Title,
		"ref":   in.Pr.Ref,
		"body":  in.Pr.Body,
	}
	return state, nil
}

func Confirm(in *api.ConfirmInput) error {
	log.Printf("Confirmed PR: %s with state: %+v\n", in.Pr.Url, in.State)
	return nil
}

func Close(in *api.CloseInput) error {
	log.Printf("Closed PR: %s with state: %+v\n", in.Pr.Url, in.State)
	return nil
}

func OpenHandler(w http.ResponseWriter, r *http.Request) {
	var input api.OpenInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	state, err := Open(&input)
	if err != nil {
		http.Error(w, "Error in Open()", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(state)
	if err != nil {
		return
	}
}

func ConfirmHandler(w http.ResponseWriter, r *http.Request) {
	var input api.ConfirmInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := Confirm(&input); err != nil {
		http.Error(w, "Error in Confirm()", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CloseHandler(w http.ResponseWriter, r *http.Request) {
	var input api.CloseInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := Close(&input); err != nil {
		http.Error(w, "Error in Close()", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
