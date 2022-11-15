package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	res := os.Getenv("TEST")
	println(res)
}
