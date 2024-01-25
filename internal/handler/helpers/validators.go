package h_help

import (
	"errors"
	"regexp"
	"users-list/server"
)

func Validate_Name(name string) error {
	regex := regexp.MustCompile(`^[A-Z][a-z]+$`)
	if !regex.MatchString(name) {
		return errors.New("invalid name")
	}
	return nil
}

func Validate_Surname(name string) error {
	regex := regexp.MustCompile(`^[A-Z][a-z]+$`)
	if !regex.MatchString(name) {
		return errors.New("invalid surname")
	}
	return nil
}

func Validate_Patronymic(name string) error {
	regex := regexp.MustCompile(`^[A-Z][a-z]+$`)
	if !regex.MatchString(name) {
		return errors.New("invalid patronymic")
	}
	return nil
}

func Validate_Post_Structure(input *server.Post_structure) error {
	if err := Validate_Name(input.Name); err != nil {
		return err
	}
	if err := Validate_Surname(input.Surname); err != nil {
		return err
	}
	if input.Patronymic != "" {
		if err := Validate_Patronymic(input.Patronymic); err != nil {
			return err
		}
	}

	return nil
}