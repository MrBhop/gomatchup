package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MrBhop/gomatchup/internal/algorithm"
	datastructures "github.com/MrBhop/gomatchup/internal/dataStructures"
)

type state struct {
	Players datastructures.Graph[string]
}

func (s *state) AddPlayers(command string, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: %s <player>\n", command)
	}
	player := args[0]

	if exists, _ := s.Players.HasNodes(player); exists {
		return fmt.Errorf("'%s' already exists", player)
	}

	s.Players.AddNode(player)
	fmt.Printf("Player '%s' added.\n", player)
	return nil
}

func (s *state) RemovePlayer(command string, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: %s <player>\n", command)
	}
	player := args[0]

	if exists, _ := s.Players.HasNodes(player); !exists {
		return fmt.Errorf("'%s' doesn't exists", player)
	}

	s.Players.RemoveNode(player)
	fmt.Printf("Player '%s' removed.\n", player)
	return nil
}

func (s *state) AddConstraint(command string, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: %s <player1> <player2>", command)
	}
	player1 := args[0]
	player2 := args[1]

	if exists, missingPlayer := s.Players.HasNodes(player1, player2); !exists {
		return fmt.Errorf("'%s' doesn't exists", missingPlayer)
	}

	s.Players.AddEdge(player1, player2)
	fmt.Printf("Added exclusion constraint, '%s' X '%s'.\n", player1, player2)
	return nil
}

func (s *state) RemoveConstraint(command string, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: %s <player1> <player2>", command)
	}
	player1 := args[0]
	player2 := args[1]

	if exists, missingPlayer := s.Players.HasNodes(player1, player2); !exists {
		return fmt.Errorf("'%s' doesn't exists", missingPlayer)
	}

	s.Players.RemoveEdge(player1, player2)
	fmt.Printf("Removed exclusion constraint, '%s' X '%s'.\n", player1, player2)
	return nil
}

func (s *state) ListPlayers(_ string, _ ...string) error {
	playerCount := s.Players.CountNodes()
	if playerCount == 0 {
		fmt.Println("Currently no players.")
		return nil
	}

	fmt.Println("Current number of players:", playerCount)
	for player := range s.Players.AllNodes().All() {
		fmt.Printf("\t- '%s'\n", player)
	}
	return nil
}

func (s *state) ListConstraints(_ string, _ ...string) error {
	constraintCount := s.Players.CountEdges()
	if constraintCount == 0 {
		fmt.Println("Currently no constraints.")
		return nil
	}

	fmt.Println("Current number of constraints:", constraintCount)
	donePlayers := datastructures.NewSet[string]()
	for player := range s.Players.AllNodes().All() {
		for neighbour := range s.Players.AdjacentNodes(player).All() {
			if donePlayers.Contains(neighbour) {
				continue
			}

			fmt.Printf("\t- '%s' X '%s'\n", player, neighbour)
		}

		donePlayers.Add(player)
	}
	return nil
}

func (s *state) GenerateTeams(command string, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: %s <number of teams>\n", command)
	}
	numberOfTeams, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Error converting 'number of teams' to number: %s", err)
	}

	teams, err := algorithm.AssignNodes(s.Players, numberOfTeams)
	if err != nil {
		if errors.Is(err, algorithm.NoValidAssignmentsError) {
			return fmt.Errorf("Couldn't find valid team assignments for the current constraints: %s", err)
		}
		return fmt.Errorf("Error assigning players: %s", err)
	}

	printTeams(teams)
	return nil
}

func printTeams(teams []datastructures.Set[string]) {
	numberOfTeams := len(teams)
	for index, team := range teams {
		fmt.Printf("Team %d of %d:\n", index + 1, numberOfTeams)
		for member := range team.All() {
			fmt.Printf("\t- '%s'\n", member)
		}
		fmt.Println()
	}
}
