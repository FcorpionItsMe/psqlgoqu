package repository

import "github.com/FcorpionItsMe/psqlgoqu/internal/domain"

type PathRepository interface {
	SavePath(path domain.PathItem) error
	GetAllPathes() ([]domain.PathItem, error)
}
