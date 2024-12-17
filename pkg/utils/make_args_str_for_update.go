package utils

import (
	"fmt"
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
)

func MakeArgsStrForUpdate(preSortedColumns []entities.Column) (string, int) {
	res := ""
	lastPlaceHolderIndex := 0
	for i := 0; i < len(preSortedColumns); i++ {
		placeHolerIndex := i + 1
		res += fmt.Sprintf("%s = $%d", preSortedColumns[i], placeHolerIndex)
		if i <= len(preSortedColumns)-2 {
			res += " ,"
		}
		lastPlaceHolderIndex = placeHolerIndex
	}
	return res, lastPlaceHolderIndex
}
