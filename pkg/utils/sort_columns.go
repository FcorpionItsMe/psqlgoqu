package utils

import (
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
)

func SortColumns(columns []entities.Column) {
	for i := 0; i < len(columns); i++ {
		for j := 1; j < len(columns); j++ {
			if columns[i].Position > columns[j].Position {
				columns[i], columns[j] = columns[j], columns[i]
			}
		}
	}
}
