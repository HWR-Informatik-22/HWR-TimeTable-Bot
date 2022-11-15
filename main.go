package main

import (
	"hwrbot/helper"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	file, err := helper.GetIcsFile("kursa")

	if err != nil {
		panic(err)
	}

	//println(file)

	/*start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(strings.NewReader(file))
	c.Start, c.End = &start, &end
	c.Parse()

	for _, e := range c.Events {
		fmt.Printf("%s on %s", e.Summary, e.Start)
	}*/

	_ = file
	println("Bot started")

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
