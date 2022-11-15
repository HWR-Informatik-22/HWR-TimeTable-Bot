package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apognu/gocal"
)

type lesson struct {
	name    string
	room    string
	teacher string
	start   *time.Time
	end     *time.Time
}

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

func ParseIcsFile(ics *gocal.Gocal) []lesson {
	array := make([]lesson, len(ics.Events))
	for key, e := range ics.Events {
		split := strings.Split(e.Summary, ";")
		array[key].name = split[1]
		array[key].teacher = split[2]
		if len(split) == 4 {
			array[key].room = split[3]
		}
		array[key].start = e.Start
		array[key].end = e.End
	}
	return array
}
