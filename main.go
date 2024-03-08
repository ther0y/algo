package main

import (
	"github.com/joho/godotenv"
	"masood.dev/algo/cmd"

	_ "masood.dev/algo/cmd/nobitex"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
