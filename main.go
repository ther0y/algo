package main

import (
	"github.com/joho/godotenv"
	"masood.dev/algo/cmd"

	_ "masood.dev/algo/cmd/code-review"
	_ "masood.dev/algo/cmd/nobitex"
	_ "masood.dev/algo/cmd/notion"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
