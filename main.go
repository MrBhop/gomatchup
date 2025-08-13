package main

import (
	"fmt"
	"log"

	"github.com/MrBhop/gomatchup/internal/repl"
)

func main() {
	replProvider := repl.NewReplCommandList(repl.ReplConfig{
		Prompt: "gomatchup >>",
	})

	fmt.Println()
	fmt.Println("Welcome to gomatchup!")

	if err := replProvider.Start(); err != nil {
		log.Fatalln(err)
	}
}
