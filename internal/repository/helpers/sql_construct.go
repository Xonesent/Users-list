package r_help

import (
	"fmt"
	"users-list/server"
)

func ExtractSQL(data *server.Patch_structure) (string, []interface{}) {
	return_value := ""
	count := 1
	var args []interface{}

	if data.Name != "" {
		return_value += fmt.Sprintf("name=$%d", count)
		count++
		args = append(args, data.Name)
	}
	if data.Surname != "" {
		if return_value != "" {
			return_value += ", "
		}
		return_value += fmt.Sprintf("surname=$%d", count)
		count++
		args = append(args, data.Surname)
	}
	if data.Patronymic != "" {
		if return_value != "" {
			return_value += ", "
		}
		return_value += fmt.Sprintf("patronymic=$%d", count)
		count++
		args = append(args, data.Patronymic)
	}
	if data.Age != 0 {
		if return_value != "" {
			return_value += ", "
		}
		return_value += fmt.Sprintf("age=$%d", count)
		count++
		args = append(args, data.Age)
	}
	if data.Gender != "" {
		if return_value != "" {
			return_value += ", "
		}
		return_value += fmt.Sprintf("gender=$%d", count)
		count++
		args = append(args, data.Gender)
	}
	if data.Nationality != "" {
		if return_value != "" {
			return_value += ", "
		}
		return_value += fmt.Sprintf("nationality=$%d", count)
		count++
		args = append(args, data.Nationality)
	}
	return_value += fmt.Sprintf(" WHERE id=$%d", count)
	args = append(args, data.Id)

	return return_value, args
}
