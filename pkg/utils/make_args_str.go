package utils

import (
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
)

func MakeArgsStr(preSortedColumns []entities.Column) string {
	res := ""
	for i := 0; i < len(preSortedColumns); i++ {
		res += preSortedColumns[i].Name
		if i <= len(preSortedColumns)-2 {
			res += ", "
		}
	}
	return res
}
