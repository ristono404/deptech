package file

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func UploadImage(encoded, key string) (string, error) {
	newFilename := fmt.Sprintf("%d-%s.jpg", time.Now().UnixNano(), key)
	file, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll("../../web/", os.ModePerm); err != nil {
		return "", err
	}

	f, err := os.Create("../../web/" + newFilename)
	if err != nil {
		return "", err
	}
	if _, err := f.Write(file); err != nil {
		return "", err
	}

	return newFilename, nil
}
