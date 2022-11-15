package helper

import (
	"fmt"
	"net/http"
	"os"
)

func GetIcsFile(course string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("BASE_ICS_URL"), course)

	file, err := http.Get(url)

	if err != nil {
		return file, err
	}

	return file, nil
}
