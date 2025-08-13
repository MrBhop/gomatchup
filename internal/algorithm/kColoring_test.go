package algorithm

import (
	"fmt"
	"math"
	"testing"

	datastructures "github.com/MrBhop/gomatchup/internal/dataStructures"
)

func TestKColoring(t *testing.T) {
	type input struct {
		players                  []string
		numberOfTeams            int
		exclusionList            [][2]string
		validAssignmentsPossible bool
	}

	createPlayerGraph := func(players []string, exclusionList [][2]string) graph[string] {
		output := datastructures.NewGraph[string]()
		for _, p := range players {
			output.AddNode(p)
		}
		for _, tuple := range exclusionList {
			output.AddEdge(tuple[0], tuple[1])
		}
		return output
	}

	checkExclusionCriteria := func(teams []set[string], constraints graph[string]) error {
		baseSize := teams[0].Count()
		for index, team := range teams {
			// check team size.
			if int(math.Abs(float64(team.Count() - baseSize))) > 1 {
				return fmt.Errorf("team sizes differ by more than 1.")
			}

			// check if the team has forbidden pairings.
			for player := range team.All() {
				// check if player has any constraints.
				if exist, _ := constraints.HasNodes(player); !exist {
					continue
				}

				for notAllowedMate := range constraints.AdjacentNodes(player).All() {
					if team.Contains(notAllowedMate) {
						return fmt.Errorf("team %d contains both '%v' and '%v'.", index, player, notAllowedMate)
					}
				}
			}
		}

		return nil
	}

	runCase := func(c input) error {
		playersGraph := createPlayerGraph(c.players, c.exclusionList)

		output, err := assignNodes(playersGraph, c.numberOfTeams)
		if err != nil {
			if !c.validAssignmentsPossible {
				return nil
			}

			return fmt.Errorf("Cases expects valid assignment, but none was found: %w", err)
		}

		// print teams
		for i, team := range output {
			t.Logf("Team %d of %d:\n", i + 1, len(output))
			counter := 1
			for player := range team.All() {
				t.Logf("%d.) %s\n", counter, player)

				counter++
			}
			t.Log()
		}

		if !c.validAssignmentsPossible {
			return fmt.Errorf("Case expected no valid assignment, but one was still found")
		}

		return checkExclusionCriteria(output, playersGraph)
	}


	cases := []input{
		// valid assignment with constraints
		{
			players: []string{
				"player1",
				"player2",
				"player3",
				"player4",
				"player5",
				"player6",
			},
			numberOfTeams: 2,
			exclusionList: [][2]string{
				{
					"player1",
					"player2",
				},
				{
					"player2",
					"player3",
				},
			},
			validAssignmentsPossible: true,
		},
		// no valid assignment - overconstrained
		{
			players: []string{
				"player1",
				"player2",
				"player3",
				"player4",
				"player5",
				"player6",
			},
			numberOfTeams: 2,
			exclusionList: [][2]string{
				{
					"player1",
					"player2",
				},
				{
					"player1",
					"player3",
				},
				{
					"player1",
					"player4",
				},
				{
					"player1",
					"player5",
				},
				{
					"player1",
					"player6",
				},
			},
			validAssignmentsPossible: false,
		},
	}

	for i, c := range cases {
		if err := runCase(c); err != nil {
			t.Errorf("case %d of %d did not pass: %v", i + 1, len(cases), err)
		}
	}
}
