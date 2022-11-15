package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetIcsFile(course string) (string, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("BASE_ICS_URL"), course)
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
