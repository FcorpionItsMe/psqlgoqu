package json_repository

import (
	"errors"
	"github.com/FcorpionItsMe/psqlgoqu/internal/domain"
	"os"
)

type Repository struct {
	pathFile   *os.File
	tablesFile *os.File
	pathes     []domain.PathItem
}

func NewRepository() (*Repository, error) {
	info, err := os.Stat("goqudata")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_ = os.Mkdir("goqudata", 0750)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	info.IsDir()
	///TODO: Пути сохранения можно вынести в конфиг
	openedPathFile, err := os.OpenFile("goqudata/pathes.json", os.O_RDWR, 0666)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = os.Create("goqudata/pathes.json")
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Cannot open path file!" + err.Error())
		}
	}
	openedTablesFile, err := os.OpenFile("goqudata/tables.json", os.O_CREATE, 0666)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = os.Create("goqudata/tables.json")
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Cannot open tables file!")
		}
	}
	return &Repository{pathFile: openedPathFile, tablesFile: openedTablesFile}, nil
}
