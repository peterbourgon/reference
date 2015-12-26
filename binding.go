package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/julienschmidt/httprouter"
)

// httpBinding provides julienschmidt/httprouter handlers for each of the
// UserService endpoints. It expects/provides all requests/responses in JSON.
type httpBinding struct {
	createEndpoint endpoint.Endpoint
	getEndpoint    endpoint.Endpoint
	deleteEndpoint endpoint.Endpoint
}

// register each method of the binding to the passed router. Keep this close
// by to the handler implementations, because the handlers will need to read
// the URL parameters we define here.
func (b httpBinding) register(r *httprouter.Router) {
	r.POST("/api/v0/users/:emailAddress", b.handleCreate)
	r.GET("/api/v0/users/:emailAddress", b.handleGet)
	r.DELETE("/api/v0/users/:emailAddress", b.handleDelete)
}

func (b httpBinding) handleCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	respondError(w, http.StatusNotImplemented, "not yet implemented")
}

func (b httpBinding) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	respondError(w, http.StatusNotImplemented, "not yet implemented")
}

func (b httpBinding) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	respondError(w, http.StatusNotImplemented, "not yet implemented")
}

// respondError in some canonical format.
func respondError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       msg,
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}
