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

type Lesson struct {
	name    string
	room    string
	teacher string
	start   string
	end     string
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

func ParseIcsFile(file string) []lesson {
	houresLeft := 24 - time.Now().Hour()
	start, end := time.Now(), time.Now().Add(time.Hour*time.Duration(houresLeft))

	ics := gocal.NewParser(strings.NewReader(file))
	ics.Start, ics.End = &start, &end
	ics.Parse()

	array := make([]lesson, len(ics.Events))

	for key, e := range ics.Events {
		split := strings.Split(e.Summary, ";")

		array[key].name = split[1]
		array[key].teacher = split[2]

		if len(split) == 4 {
			array[key].room = split[3]
		}

		array[key].start = e.Start.Format("15:04")
		array[key].end = e.End.Format("15:04")
	}
	return array
}
