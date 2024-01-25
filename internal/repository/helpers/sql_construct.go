package r_help

import (
	"fmt"
	"users-list/server"
)

func Extract_SQL_Get(data *server.Get_structure) (string, []interface{}) {
	return_value := ""
	count := 1
	var args []interface{}

	if data.Id != 0 {
		return_value += fmt.Sprintf("id=$%d", count)
		count++
		args = append(args, data.Id)
	}
	if data.Name != "" {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("name=$%d", count)
		count++
		args = append(args, data.Name)
	}
	if data.Surname != "" {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("surname=$%d", count)
		count++
		args = append(args, data.Surname)
	}
	if data.Patronymic != "" {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("patronymic=$%d", count)
		count++
		args = append(args, data.Patronymic)
	}
	if data.Age != 0 {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("age=$%d", count)
		count++
		args = append(args, data.Age)
	}
	if data.Gender != "" {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("gender=$%d", count)
		count++
		args = append(args, data.Gender)
	}
	if data.Nationality != "" {
		if return_value != "" {
			return_value += " AND "
		}
		return_value += fmt.Sprintf("nationality=$%d", count)
		count++
		args = append(args, data.Nationality)
	}
	if data.Limit != 0 {
		return_value += fmt.Sprintf(" LIMIT $%d", count)
		count++
		args = append(args, data.Limit)
	}

	if data.Offset != 0 {
		return_value += fmt.Sprintf(" OFFSET $%d", count)
		args = append(args, data.Offset)
	}

	return return_value, args
}

func Extract_SQL_Patch(data *server.Patch_structure) (string, []interface{}) {
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
