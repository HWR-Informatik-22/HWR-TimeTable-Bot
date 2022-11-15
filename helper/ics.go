package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetIcsFile(course string) (io.ReadCloser, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("BASE_ICS_URL"), course)

	file, err := http.Get(url)

	if err != nil {
		return file.Request.Body, err
	}

	return file.Request.Body, nil
}
