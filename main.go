package main

import (
	"fmt"
	"hwrbot/helper"

	"github.com/apognu/gocal"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	file, err := helper.GetIcsFile("kursa")

	if err != nil {
		panic(err)
	}

	//start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(file)
	//c.Start, c.End = &start, &end
	c.Parse()

	for _, e := range c.Events {
		fmt.Printf("%s on %s by %s", e.Summary, e.Start, e.Organizer.Cn)
	}

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
