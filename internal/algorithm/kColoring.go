package algorithm

import (
	"errors"

	"github.com/MrBhop/gomatchup/internal/dataStructures"
)

type constraints = map[string]Set[string]

func playerCanBeAssignedToTeam(player string, team Set[string], exclusionMap map[string]Set[string]) bool {
	for p := range team {
		if exclusionMap[p].Contains(player) {
			return false
		}
	}

	return true
}

func assignPlayers(players []string, exclusionMap constraints, numberOfTeams int) ([]Set[string], error) {
	playerStack := datastructures.NewSimpleStack(players)
	teams := make([]Set[string], numberOfTeams)
	for i := range teams {
		teams[i] = Set[string]{}
	}
	maxTeamSize := len(players) / numberOfTeams
	if len(players) % numberOfTeams != 0 {
		maxTeamSize += 1
	}

	output := assignPlayersR(playerStack, teams, exclusionMap, maxTeamSize)
	if output == nil {
		return nil, errors.New("could not find valid assignment.")
	}

	return output, nil
}

func assignPlayersR(players datastructures.SimpleStack[string], teams []Set[string], exclusionMap constraints, maxTeamSize int) []Set[string] {
	player, exists := players.Pop()
	if !exists {
		return teams
	}

	for index, team := range teams {
		if !playerCanBeAssignedToTeam(player, team, exclusionMap) {
			continue
		}

		if len(team) >= maxTeamSize {
			continue
		}

		teams[index].Add(player)

		if players.IsEmpty() {
			return teams
		}

		newTeams := assignPlayersR(players, teams, exclusionMap, maxTeamSize)
		if newTeams != nil {
			return newTeams
		}

		teams[index].Remove(player)
		players.Push()
	}

	return nil
}
