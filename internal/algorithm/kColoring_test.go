package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func TestKColoring(t *testing.T) {
	type input struct {
		players                  []string
		numberOfTeams            int
		exclusionList             [][2]string
		validAssignmentsPossible bool
	}

	createExclusionMap := func(exclusionList [][2]string) map[string]Set[string] {
		output := map[string]Set[string]{}
		for _, tuple := range exclusionList {
			val1, val2 := tuple[0], tuple[1]

			if _, exists := output[val1]; !exists {
				output[val1] = Set[string]{}
			}
			if _, exists := output[val2]; !exists {
				output[val2] = Set[string]{}
			}

			i := output[val1]
			i.Add(val2)

			j := output[val2]
			j.Add(val1)
		}
		return output
	}

	checkExclusionCriteria := func(teams []Set[string], exclusionMap constraints) error {
		// check team sizes.
		baseSize := len(teams[0])
		for i := 1; i < len(teams); i++ {
			if int(math.Abs(float64(len(teams[i]) - baseSize))) > 1 {
				return fmt.Errorf("team sizes differ by more than 1.")
			}
		}

		// check against exclusion map.
		for i, team := range teams {
			for player := range team {
				for player2 := range team {
					m, exists := exclusionMap[player]
					if !exists {
						continue
					}

					if m.Contains(player2) {
						return fmt.Errorf("team %d violates exclusion list. Contains %s & %s.", i + 1, player, player2)
					}
				}
			}
		}

		return nil
	}

	runCase := func(c input) error {
		exclusionMap := createExclusionMap(c.exclusionList)

		output, err := assignPlayers(c.players, exclusionMap, c.numberOfTeams)
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
			for player := range team {
				t.Logf("%d.) %s\n", counter, player)

				counter++
			}
			t.Log()
		}

		if !c.validAssignmentsPossible {
			return fmt.Errorf("Case expected no valid assignment, but one was still found")
		}

		return checkExclusionCriteria(output, exclusionMap)
	}


	cases := []input{
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
