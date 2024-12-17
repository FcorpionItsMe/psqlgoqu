package domain

import "github.com/FcorpionItsMe/psqlgoqu/pkg/entities"

type TableItem struct {
	PathParent PathItem
	Name       string
	Columns    []entities.Column
}
