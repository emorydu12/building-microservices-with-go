package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emorydu/building-microservices-with-go/data"
)

type searchRequest struct {
	Query string `json:"query"`
}

type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}

type Search struct {
	DataStore data.Store
}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	req := new(searchRequest)
	err := decoder.Decode(req)
	if err != nil || len(req.Query) < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	kittens := s.DataStore.Search(req.Query)
	response := &searchResponse{Kittens: kittens}

	_ = json.NewEncoder(w).Encode(response)
}
