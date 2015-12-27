package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

func bindingFixtures(t *testing.T) (UserService, http.Handler) {
	r := httprouter.New()
	s := inmemFixture(t)
	httpBinding{
		ctx:            context.Background(),
		createEndpoint: makeCreateEndpoint(s),
		getEndpoint:    makeGetEndpoint(s),
		deleteEndpoint: makeDeleteEndpoint(s),
	}.register(r)
	return s, r
}

func TestBindingCreate(t *testing.T) {
	s, h := bindingFixtures(t)
	server := httptest.NewServer(h)
	defer server.Close()

	// Create a new user
	const (
		emailAddress = "test@email.biz"
		realName     = "foobar"
		password     = "barbaz"
	)
	resp, err := http.Post(fmt.Sprintf(
		"%s/api/v0/users/%s?RealName=%s&Password=%s",
		server.URL,
		emailAddress,
		realName,
		password,
	), "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
	resp.Body.Close()

	// Get the created user
	user, err := s.Get(emailAddress)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := realName, user.RealName; want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}

func TestBindingGet(t *testing.T) {
	_, h := bindingFixtures(t)
	server := httptest.NewServer(h)
	defer server.Close()

	// Get the pre-existing user
	resp, err := http.Get(fmt.Sprintf(
		"%s/api/v0/users/%s",
		server.URL,
		emailAddress,
	))
	if err != nil {
		t.Fatal(err)
	}
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
	defer resp.Body.Close()

	// It should exist
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Error(err)
	}
	if want, have := realName, user.RealName; want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}

func TestBindingDelete(t *testing.T) {
	s, h := bindingFixtures(t)
	server := httptest.NewServer(h)
	defer server.Close()

	// Delete the pre-existing user
	req, err := http.NewRequest("DELETE", fmt.Sprintf(
		"%s/api/v0/users/%s",
		server.URL,
		emailAddress,
	), nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
	defer resp.Body.Close()

	// It should no longer exist
	_, err = s.Get(emailAddress)
	if want, have := ErrEmailAddressNotFound, err; want != have {
		t.Errorf("want %v, have %v", want, have)
	}
}
