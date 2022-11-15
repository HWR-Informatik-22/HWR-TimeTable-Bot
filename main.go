package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	println("started bot")
	res := os.Getenv("TEST")
	println(res)
}
