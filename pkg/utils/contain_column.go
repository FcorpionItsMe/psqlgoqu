package utils

import (
	"github.com/FcorpionItsMe/psqlgoqu/pkg/entities"
)

func ContainColumn(t entities.Table, clmn entities.Column) bool {
	colmns := t.GetColumns()
	for _, e := range colmns {
		if e.Name == clmn.Name && e.Position == clmn.Position {
			return true
		}
	}
	return false
}
