package handlers

import "net/http"

type bangHandler struct {

}

func (b *bangHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("Something gone wrong again")
}

func NewBangHandler() http.Handler {
	return &bangHandler{}
}