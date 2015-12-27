package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

// CreateRequest is for the corresponding endpoint.
type CreateRequest struct {
	EmailAddress      string
	RealName          string
	PlaintextPassword string
}

// CreateResponse is for the corresponding endpoint.
type CreateResponse struct {
	//
}

func makeCreateEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		err := s.Create(req.EmailAddress, req.RealName, req.PlaintextPassword)
		return CreateResponse{}, err
	}
}

// GetRequest is for the corresponding endpoint.
type GetRequest struct {
	EmailAddress string
}

// GetResponse is for the corresponding endpoint.
type GetResponse struct {
	User User
}

func makeGetEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		u, err := s.Get(req.EmailAddress)
		return GetResponse{User: u}, err
	}
}

// DeleteRequest is for the corresponding endpoint.
type DeleteRequest struct {
	EmailAddress string
}

// DeleteResponse is for the corresponding endpoint.
type DeleteResponse struct {
	//
}

func makeDeleteEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(req.EmailAddress)
		return DeleteResponse{}, err
	}
}
