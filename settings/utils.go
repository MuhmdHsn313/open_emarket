package settings

import (
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetImagePath(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	return MEDIA_PATH + uuid.New().String() + ext
}
