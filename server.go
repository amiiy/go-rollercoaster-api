package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	Name         string
	Manufacturer string
	ID           string
	InPark       string
	Height       int
}
type coasterHandler struct {
	store map[string]Coaster
}

func (h *coasterHandler) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method no allowed"))
		return

	}
}
func (h *coasterHandler) post(w http.ResponseWriter, r *http.Request) {

}

func (h *coasterHandler) get(w http.ResponseWriter, r *http.Request) {

	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newCoasterHandler() *coasterHandler {
	return &coasterHandler{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:         "Fury 222",
				Height:       99,
				ID:           "id1",
				InPark:       "sawq",
				Manufacturer: "B+M",
			},
		},
	}
}
func main() {
	coasterHandler := newCoasterHandler()
	http.HandleFunc("/coaster", coasterHandler.coasters)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
