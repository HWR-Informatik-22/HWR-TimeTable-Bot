package helper

import (
	"fmt"
	"os"
)

func GetIcsFile(course string) {
	url := fmt.Sprintf("%s%s", os.Getenv("BASE_ICS_URL"), course)

	println(url)
}
