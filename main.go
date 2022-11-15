package main

import (
	"fmt"
	"hwrbot/helper"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	file, err := helper.GetIcsFile("kursa")

	if err != nil {
		panic(err)
	}

	//println(file)

	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(strings.NewReader(file))
	c.Start, c.End = &start, &end
	c.Parse()

	type lesson struct {
		name    string
		room    string
		teacher string
		start   *time.Time
		end     *time.Time
	}

	//array := make([]lesson, len(c.Events))
	for _, e := range c.Events {
		fmt.Println("%v", strings.SplitN(e.Summary, ";", 4))
		/*
			array[key].name = strings.SplitN(e.Summary, ";", 4)[1]
			array[key].teacher = strings.SplitN(e.Summary, ";", 4)[3]
			array[key].start = e.Start
			array[key].end = e.End
		*/

	}
	//fmt.Printf("%v", array[0])
	/*
	   ics, err := http.Get(os.Getenv("BASE_ICS_URL"))

	   	if err != nil {
	   		panic(err)
	   	}

	   _ = ics

	   println("started bot")
	   res := os.Getenv("BASE_ICS_URL")
	   println(res)
	*/
}
