package utils

import (
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
)

func WithoutColumn(clmns []entities.Column, c entities.Column) []entities.Column {
	var res []entities.Column
	for i := 0; i < len(clmns); i++ {
		if clmns[i].Equals(c) {
			continue
		}
		res = append(res, clmns[i])
	}
	return res
}
