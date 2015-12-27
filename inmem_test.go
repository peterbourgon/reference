package main

import "testing"

const (
	emailAddress      = "foo@bar.baz"
	realName          = "Foo Bar"
	plaintextPassword = "hunter2"
)

func inmemFixture(t *testing.T) UserService {
	s := newInmemUserService()
	if err := s.Create(emailAddress, realName, plaintextPassword); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestInmemCreate(t *testing.T) {
	s := inmemFixture(t)
	if want, have := error(nil), s.Create("alterate@email.address", "Bar Baz", "abcdefg"); want != have {
		t.Errorf("want %v, have %v", want, have)
	}
	if want, have := ErrEmailAddressAlreadyExists, s.Create(emailAddress, realName, plaintextPassword); want != have {
		t.Errorf("want %v, have %v", want, have)
	}
}

func TestInmemGet(t *testing.T) {
	s := inmemFixture(t)
	u, err := s.Get(emailAddress)
	if err != nil {
		t.Error(err)
	}
	if want, have := emailAddress, u.EmailAddress; want != have {
		t.Errorf("want %q, have %q", want, have)
	}
	if want, have := realName, u.RealName; want != have {
		t.Errorf("want %q, have %q", want, have)
	}
}

func TestInmemDelete(t *testing.T) {
	s := inmemFixture(t)
	if want, have := error(nil), s.Delete(emailAddress); want != have {
		t.Errorf("want %v, have %v", want, have)
	}
	if want, have := ErrEmailAddressNotFound, s.Delete(emailAddress); want != have {
		t.Errorf("want %v, have %v", want, have)
	}
}
