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
	houresLeft := 24 - time.Now().Hour()
	start, end := time.Now(), time.Now().Add(time.Hour*time.Duration(houresLeft))

	c := gocal.NewParser(strings.NewReader(file))
	c.Start, c.End = &start, &end
	c.Parse()

	fmt.Print(helper.ParseIcsFile(c))
}
