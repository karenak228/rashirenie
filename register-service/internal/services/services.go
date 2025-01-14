package services //бизнес логика

import (
	"errors"
	"rashirenie/register-service/internal/repository"
	"rashirenie/register-service/internal/requests"
)

type Services struct {
	repository *repository.Repositori
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

func (s *Services) Registration(r requests.Client) error {
	if r.Email == "a@a" {
		return nil
	}
	return errors.New(" ")
}
