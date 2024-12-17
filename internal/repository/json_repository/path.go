package json_repository

import (
	"encoding/json"
	"github.com/FcorpionItsMe/psqlgoqu/internal/domain"
	"github.com/FcorpionItsMe/psqlgoqu/internal/utils/files"
	"os"
	"strings"
)

func (r *Repository) SavePath(path domain.PathItem) error {
	if r.pathes == nil {
		loadedPathes, err := r.GetAllPathes()
		if err != nil {
			return err
		}
		r.pathes = loadedPathes
	}
	r.pathes = append(r.pathes, path)
	err := r.pathFile.Close()
	if err != nil {
		return err
	}
	r.pathFile, err = os.Create("goqudata/pathes.json")
	if err != nil {
		return err
	}
	var splittedPathesStr []string
	for i := 0; i < len(r.pathes); i++ {
		jsonPath, err := json.Marshal(r.pathes[i])
		if err != nil {
			return err
		}
		splittedPathesStr = append(splittedPathesStr, string(jsonPath))
	}
	resToSave := strings.Join(splittedPathesStr, "\n")
	_, err = r.pathFile.WriteString(resToSave)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetAllPathes() ([]domain.PathItem, error) {
	pathFileContent, err := files.ReadContent(r.pathFile)
	if err != nil {
		return nil, err
	}
	splittedFileContent := strings.Split(string(pathFileContent), "\n")
	if len(splittedFileContent) == 1 && splittedFileContent[0] == "" {
		return []domain.PathItem{}, nil
	}
	res := make([]domain.PathItem, len(splittedFileContent))
	for i := 0; i < len(splittedFileContent); i++ {
		var pathFromJson domain.PathItem
		err = json.Unmarshal([]byte(splittedFileContent[i]), &pathFromJson)
		if err != nil {
			return nil, err
		}
		res[i] = pathFromJson
	}
	if r.pathes == nil {
		r.pathes = res
	}
	return res, nil
}

func (r *Repository) FindPathItemByActualPath(path string) (domain.PathItem, bool) {
	for i := 0; i < len(r.pathes); i++ {
		if r.pathes[i].ActualPath == path {
			return r.pathes[i], true
		}
	}
	return domain.PathItem{}, false
}
func (r *Repository) FindPathItemByFirstPath(path string) (domain.PathItem, bool) {
	for i := 0; i < len(r.pathes); i++ {
		if r.pathes[i].FirstCallPath == path {
			return r.pathes[i], true
		}
	}
	return domain.PathItem{}, false
}
func (r *Repository) FindPathItemByLastPath(path string) (domain.PathItem, bool) {
	for i := 0; i < len(r.pathes); i++ {
		if r.pathes[i].LastCallPath == path {
			return r.pathes[i], true
		}
	}
	return domain.PathItem{}, false
}
