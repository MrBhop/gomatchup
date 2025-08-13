package algorithm

import (
	"fmt"

	datastructures "github.com/MrBhop/gomatchup/internal/dataStructures"
)

func playerCanBeAssignedToTeam[T comparable](player T, team set[T], constraints graph[T]) bool {
	for teamMember := range team.All() {
		if constraints.HasEdge(player, teamMember) {
			return false
		}
	}

	return true
}

func newSliceOfSets[T comparable](length int) []set[T] {
	output := make([]set[T], length)
	for i := range output {
		output[i] = datastructures.NewSet[T]()
	}

	return output
}

func getMaxTeamSize(numberOfPlayers, numberOfTeams int) int {
	n := numberOfPlayers / numberOfTeams
	if numberOfPlayers % numberOfTeams != 0 {
		n++
	}
	return n
}

func noValidAssignmentsError() error {
	return fmt.Errorf("could not find valid assignments.")
}

func assignPlayers[T comparable](players graph[T], numberOfTeams int) ([]set[T], error) {
	maxTeamSize := getMaxTeamSize(players.CountNodes(), numberOfTeams)
	teams := newSliceOfSets[T](numberOfTeams)

	// assign players with constraints
	playerStack := datastructures.NewSimpleStack(players.ConnectedNodes())
	teams = assignPlayersR(playerStack, teams, players, maxTeamSize)
	if teams == nil {
		return nil, noValidAssignmentsError()
	}

	// assign players without constraints.
	nextTeamToAdd := 0
	for player := range players.UnconnectedNodes().All() {
		for teams[nextTeamToAdd].Count() >= maxTeamSize {
			if nextTeamToAdd >= len(teams) - 1 {
				return nil, fmt.Errorf("Error assigning unconstrained players.")
			}
			nextTeamToAdd++
		}

		teams[nextTeamToAdd].Add(player)
	}

	return teams, nil
}

func assignPlayersR[T comparable](players simpleStack[T], teams[]set[T], constraints graph[T], maxTeamSize int) []set[T] {
	currentPlayer, exists := players.Pop()
	if !exists {
		return teams
	}
	
	for _, team := range teams {
		if team.Count() >= maxTeamSize {
			continue
		}

		if !playerCanBeAssignedToTeam(currentPlayer, team, constraints) {
			continue
		}

		team.Add(currentPlayer)

		if newTeams := assignPlayersR(players, teams, constraints, maxTeamSize); newTeams != nil {
			return newTeams
		}

		team.Remove(currentPlayer)

		if players.IsEmpty() {
			return teams
		}
	}

	players.Push()
	return nil
}
