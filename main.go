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

	println(file)

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
