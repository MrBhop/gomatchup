package algorithm

import (
	"errors"
	"fmt"

	datastructures "github.com/MrBhop/gomatchup/internal/dataStructures"
)

var NoValidAssignmentsError = errors.New("could not find valid assignments.")

func nodeCanBeAssignedToGroup[T comparable](node T, group set[T], constraints graph[T]) bool {
	for groupMember := range group.All() {
		if constraints.HasEdge(node, groupMember) {
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

func getMaxGroupSize(numberOfNodes, numberOfGroups int) int {
	n := numberOfNodes / numberOfGroups
	if numberOfNodes % numberOfGroups != 0 {
		n++
	}
	return n
}

func AssignNodes[T comparable](nodes datastructures.Graph[T], numberOfGroups int) ([]datastructures.Set[T], error) {
	maxGroupSize := getMaxGroupSize(nodes.CountNodes(), numberOfGroups)
	groups := newSliceOfSets[T](numberOfGroups)

	// assign nodes with constraints
	nodeStack := datastructures.NewSimpleStack(nodes.ConnectedNodes())
	groups = assignNodesR(nodeStack, groups, nodes, maxGroupSize)
	if groups == nil {
		return nil, NoValidAssignmentsError
	}

	// assign nodes without constraints.
	nextGroupToAddIndex := 0
	for node := range nodes.UnconnectedNodes().All() {
		for groups[nextGroupToAddIndex].Count() >= maxGroupSize {
			if nextGroupToAddIndex >= len(groups) - 1 {
				return nil, fmt.Errorf("Error assigning unconstrained node.")
			}
			nextGroupToAddIndex++
		}

		groups[nextGroupToAddIndex].Add(node)
	}

	return groups, nil
}

func assignNodesR[T comparable](nodes simpleStack[T], groups[]set[T], constraints graph[T], maxGroupSize int) []set[T] {
	currentNode, exists := nodes.Pop()
	if !exists {
		return groups
	}
	
	for _, group := range groups {
		if group.Count() >= maxGroupSize {
			continue
		}

		if !nodeCanBeAssignedToGroup(currentNode, group, constraints) {
			continue
		}

		group.Add(currentNode)

		if newGroups := assignNodesR(nodes, groups, constraints, maxGroupSize); newGroups != nil {
			return newGroups
		}

		group.Remove(currentNode)

		if nodes.IsEmpty() {
			return groups
		}
	}

	nodes.Push()
	return nil
}
