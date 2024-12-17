package goqu_errors

import (
	"fmt"
)

type TableDoesntContainColumn struct {
	T string
	C string
}

func (err TableDoesntContainColumn) Error() string {
	return fmt.Sprintf("Table \"%s\" does not contain column \"%s\"", err.T, err.C)
}
