package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

// httpBinding provides julienschmidt/httprouter handlers for each of the
// UserService endpoints. It expects/provides all requests/responses in JSON.
type httpBinding struct {
	ctx            context.Context
	createEndpoint endpoint.Endpoint
	getEndpoint    endpoint.Endpoint
	deleteEndpoint endpoint.Endpoint
}

// register each method of the binding to the passed router. Keep this close
// by to the handler implementations, because the handlers will need to read
// the URL parameters we define here.
func (b httpBinding) register(r *httprouter.Router) {
	r.POST("/api/v0/users/:EmailAddress", b.handleCreate)
	r.GET("/api/v0/users/:EmailAddress", b.handleGet)
	r.DELETE("/api/v0/users/:EmailAddress", b.handleDelete)
}

func (b httpBinding) handleCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	emailAddress := p.ByName("EmailAddress")
	if emailAddress == "" {
		respondError(w, http.StatusBadRequest, ErrEmailAddressNotProvided)
		return
	}

	realName := r.FormValue("RealName")
	plaintextPassword := r.FormValue("Password")
	if plaintextPassword == "" {
		respondError(w, http.StatusBadRequest, ErrBadPassword)
		return
	}

	if _, err := b.createEndpoint(b.ctx, CreateRequest{
		EmailAddress:      emailAddress,
		RealName:          realName,
		PlaintextPassword: plaintextPassword,
	}); err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondSuccess(w, nil)
}

func (b httpBinding) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	emailAddress := p.ByName("EmailAddress")
	if emailAddress == "" {
		respondError(w, http.StatusBadRequest, ErrEmailAddressNotProvided)
		return
	}

	response, err := b.getEndpoint(b.ctx, GetRequest{
		EmailAddress: emailAddress,
	})
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondSuccess(w, response.(GetResponse).User)
}

func (b httpBinding) handleDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	emailAddress := p.ByName("EmailAddress")
	if emailAddress == "" {
		respondError(w, http.StatusBadRequest, ErrEmailAddressNotProvided)
		return
	}

	if _, err := b.deleteEndpoint(b.ctx, DeleteRequest{
		EmailAddress: emailAddress,
	}); err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondSuccess(w, nil)
}

// respondError in some canonical format.
func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err,
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}

// respondSuccess in some canonical format.
func respondSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
