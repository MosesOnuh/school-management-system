package error_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ParseErr(err error) AppErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return AppNotFoundError("no record found that matches the given id")
		}
		return AppInternalServerError(fmt.Sprintf("error when trying to save entity: %s", err.Error()))
	}
	switch sqlErr.Number {
	case 1062:
		return AppInternalServerError("title already in use")
	}
	return AppInternalServerError(fmt.Sprintf("error when processing request: %s", err.Error()))
}