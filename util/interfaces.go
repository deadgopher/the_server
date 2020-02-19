package util

import "database/sql"

type Model interface {
	Type()string
	Connect(*sql.DB)
	Save()error
}
