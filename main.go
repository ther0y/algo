package main

import (
	"github.com/joho/godotenv"
	"masood.dev/algo/cmd"

	_ "masood.dev/algo/cmd/clickup"
	_ "masood.dev/algo/cmd/code-review"
	_ "masood.dev/algo/cmd/gitlab"
	_ "masood.dev/algo/cmd/nobitex"
	_ "masood.dev/algo/cmd/notion"
	_ "masood.dev/algo/cmd/vercel"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
