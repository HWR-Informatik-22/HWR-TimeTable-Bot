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
	fmt.Print(helper.ParseIcsFile(file))
	fmt.Scanln()
}
