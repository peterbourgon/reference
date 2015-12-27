package main

import (
	"sync"
	"time"
)

type inmemUserService struct {
	mtx   sync.Mutex
	users map[string]inmemUser
}

type inmemUser struct {
	User
	salt           string
	saltedPassword string
}

func newInmemUserService() UserService {
	return &inmemUserService{
		users: map[string]inmemUser{},
	}
}

func (s *inmemUserService) Create(emailAddress, realName, plaintextPassword string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.users[emailAddress]; ok {
		return ErrEmailAddressAlreadyExists
	}

	salt, digest := salt(plaintextPassword)
	s.users[emailAddress] = inmemUser{
		User: User{
			EmailAddress: emailAddress,
			RealName:     realName,
			CreatedAt:    time.Now().UTC(),
			UpdatedAt:    time.Now().UTC(),
		},
		salt:           salt,
		saltedPassword: digest,
	}
	return nil
}

func (s *inmemUserService) Get(emailAddress string) (User, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	u, ok := s.users[emailAddress]
	if !ok {
		return User{}, ErrEmailAddressNotFound
	}

	return u.User, nil
}

func (s *inmemUserService) Delete(emailAddress string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.users[emailAddress]; !ok {
		return ErrEmailAddressNotFound
	}

	delete(s.users, emailAddress)
	return nil
}
