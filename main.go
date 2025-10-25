package main

import (
	"fmt"
	"log"

	datastructures "github.com/MrBhop/gomatchup/internal/dataStructures"
	repl "github.com/MrBhop/goReplUtils"
)

func main() {
	programState := &state{
		Players: datastructures.NewGraph[string](),
	}

	replProvider := repl.NewReplCommandList(repl.ReplConfig{
		Prompt: "gomatchup >>",
	})

	replProvider.AddCommand("add-player", "Add a player to the pool.", programState.AddPlayers)
	replProvider.AddCommand("remove-player", "Remove a player from the pool.", programState.RemovePlayer)
	replProvider.AddCommand("add-constraint", "Add an exclusion constraint.", programState.AddConstraint)
	replProvider.AddCommand("remove-constraint", "Remove an exclusion constraint.", programState.RemoveConstraint)
	replProvider.AddCommand("list-players", "List current players in the pool.", programState.ListPlayers)
	replProvider.AddCommand("list-constraints", "List current exclusion constraints.", programState.ListConstraints)
	replProvider.AddCommand("generate-teams", "Generate teams based on the current players and constraints.", programState.GenerateTeams)

	fmt.Println()
	fmt.Println("Welcome to gomatchup!")

	if err := replProvider.Start(); err != nil {
		log.Fatalln(err)
	}
}
