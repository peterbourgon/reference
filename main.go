package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"

	"github.com/go-kit/kit/log"
)

func main() {

	// Business domain
	var s UserService
	{
		s = newInmemUserService()
	}

	// Transport domain
	r := httprouter.New()
	httpBinding{
		// This is incredibly laborious when we want to add e.g. rate
		// limiters. It would be better to bundle all the endpoints up,
		// somehow... or, use code generation, of course.
		createEndpoint: makeCreateEndpoint(s),
		getEndpoint:    makeGetEndpoint(s),
		deleteEndpoint: makeDeleteEndpoint(s),
	}.register(r)

	// Mechanical stuff
	logger := log.NewJSONLogger(os.Stdout)
	errc := make(chan error)

	go func() {
		logger.Log("msg", "HTTP server listening on :8080")
		errc <- http.ListenAndServe(":8080", r)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logger.Log("signal", fmt.Sprint(<-c))
		errc <- nil
	}()

	if err := <-errc; err != nil {
		os.Exit(1)
	}
}
