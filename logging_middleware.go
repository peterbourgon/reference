package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

// loggingMiddleware wraps UserService and logs each call.
type loggingMiddleware struct {
	log.Logger
	UserService
}

func (mw loggingMiddleware) Create(emailAddress, realName, plaintextPassword string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "create",
			"email_address", emailAddress,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return mw.UserService.Create(emailAddress, realName, plaintextPassword)
}

func (mw loggingMiddleware) Get(emailAddress string) (u User, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "get",
			"email_address", emailAddress,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return mw.UserService.Get(emailAddress)
}

func (mw loggingMiddleware) Delete(emailAddress string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "delete",
			"email_address", emailAddress,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	return mw.UserService.Delete(emailAddress)
}
