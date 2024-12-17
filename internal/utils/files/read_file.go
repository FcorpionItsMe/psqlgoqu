package files

import "os"

func ReadContent(file *os.File) ([]byte, error) {
	filestat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileContent := make([]byte, filestat.Size())
	_, err = file.Read(fileContent)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}
