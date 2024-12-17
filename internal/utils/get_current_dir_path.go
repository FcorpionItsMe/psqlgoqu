package utils

import (
	"errors"
	"log"
	"os"
)

func GetCurrentDirPath() (string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", errors.New("cannot find root dir")
	}
	return currDir, nil
}
