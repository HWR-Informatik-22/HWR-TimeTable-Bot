package main

import (
	"fmt"
	"hwrbot/helper"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	file, err := helper.GetIcsFile("kursa")

	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	fmt.Print(helper.ParseIcsFile(file))
=======

	//println(file)
	houresLeft := 24 - time.Now().Hour()
	start, end := time.Now(), time.Now().Add(time.Hour*time.Duration(houresLeft))

	c := gocal.NewParser(strings.NewReader(file))
	c.Start, c.End = &start, &end
	c.Parse()

	fmt.Print(helper.ParseIcsFile(c))

	fmt.Scanln()
>>>>>>> 98e87a366f8ae1eced6ed0f47a0647946f1f7d3c
}
