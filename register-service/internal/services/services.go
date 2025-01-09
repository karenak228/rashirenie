package services //бизнес логика

import (
	"errors"
	"rashirenie/register-service/internal/requests"
)

type Services struct {
}

func New() *Services {
	return &Services{}
}

func (s *Services) Login(r requests.Login) error {
	if r.Email == "a@a" {
		return nil
	}
	return errors.New(" ")
}
