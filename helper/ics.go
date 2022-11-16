package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	ical "github.com/arran4/golang-ical"
)

type Lesson struct {
	name    string
	room    string
	teacher string
	start   string
	end     string
}

func GetIcsContent(course string) (string, error) {
	url := os.Getenv("BASE_ICS_URL") + course

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

func ParseEventDetails(event *ical.VEvent) (l *Lesson, e error) {

	eventStart, err := event.GetStartAt()
	if err != nil {
		return nil, err
	}

	property := event.GetProperty(ical.ComponentPropertySummary)
	if property == nil {
		return nil, fmt.Errorf("Missing summary property")
	}

	summary := property.Value
	split := strings.Split(summary, "\\;")

	eventEnd, err := event.GetEndAt()
	if err != nil {
		return nil, err
	}

	les := Lesson{
		name:    split[1],
		teacher: split[2],
		room:    split[3],
		start:   eventStart.Format("15:04"),
		end:     eventEnd.Format("15:04"),
	}

	if strings.Contains(summary, "ONLINE") {
		les.room = "ONLINE"
	}

	return &les, nil
}

func ParseIcsContent(content string) (*[]Lesson, error) {
	houresLeft := 24 - time.Now().Hour()
	start, end := time.Now(), time.Now().Add(time.Hour*time.Duration(houresLeft))

	cal, err := ical.ParseCalendar(strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	array := make([]Lesson, 0)

	for _, event := range cal.Events() {

		eventStart, err := event.GetStartAt()
		if err != nil {
			continue
		}

		if eventStart.Before(start) || eventStart.After(end) {
			continue
		}

		l, err := ParseEventDetails(event)
		if err != nil {
			fmt.Println("An error occured while parsing the event details: ", err)
			continue
		}

		array = append(array, *l)
	}

	return &array, nil
}
