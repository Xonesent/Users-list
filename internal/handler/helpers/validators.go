package h_help

import (
	"errors"
	"regexp"
	"users-list/server"
)

func Validate_Id(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return nil
}

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

func Validate_Age(age int) error {
	if age > 150 {
		return errors.New("invalid age")
	}
	return nil
}

func Validate_Gender(gender string) error {
	if gender != "male" && gender != "female" {
		return errors.New("invalid gender")
	}
	return nil
}

func Validate_Nationality(nationality string) error {
	regex := regexp.MustCompile(`^[A-Z]{2}$`)
	if !regex.MatchString(nationality) {
		return errors.New("invalid nationality")
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

func Validate_Patch_Structure(input *server.Patch_structure) error {
	count := 0
	if err := Validate_Id(input.Id); err != nil {
		return err
	}

	if input.Name != "" {
		if err := Validate_Name(input.Name); err != nil {
			return err
		}
		count++
	}

	if input.Surname != "" {
		if err := Validate_Surname(input.Surname); err != nil {
			return err
		}
		count++
	}

	if input.Patronymic != "" {
		if err := Validate_Patronymic(input.Patronymic); err != nil {
			return err
		}
		count++
	}

	if input.Age != 0 {
		if err := Validate_Age(input.Age); err != nil {
			return err
		}
		count++
	}

	if input.Gender != "" {
		if err := Validate_Gender(input.Gender); err != nil {
			return err
		}
		count++
	}

	if input.Nationality != "" {
		if err := Validate_Nationality(input.Nationality); err != nil {
			return err
		}
		count++
	}

	if count == 0 {
		return errors.New("nothing to patch")
	}
	return nil
}