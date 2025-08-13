package datastructures

type Graph[T comparable] interface {
	AddNode(T)
	RemoveNode(T)
	HasNodes(...T) (bool, T)
	CountNodes() int
	AddEdge(T, T)
	RemoveEdge(T, T)
	HasEdge(T, T) bool
	CountEdges() int
	AdjacentNodes(T) Set[T]
	AllNodes() Set[T]
	ConnectedNodes() Set[T]
	UnconnectedNodes() Set[T]
}

type graphConcrete[T comparable] map[T]Set[T]

func NewGraph[T comparable]() Graph[T] {
	return graphConcrete[T]{}
}

func (g graphConcrete[T]) AddNode(node T) {
	if exist, _ := g.HasNodes(node); exist {
		return
	}

	g[node] = NewSet[T]()
}

func (g graphConcrete[T]) RemoveNode(node T) {
	if exist, _ := g.HasNodes(node); !exist {
		return
	}

	// remove connections to the target node.
	for neighbour := range g.AdjacentNodes(node).All() {
		g[neighbour].Remove(node)
	}

	delete(g, node)
}

func (g graphConcrete[T]) HasNodes(nodes ...T) (exists bool, missingNode T) {
	for _, n := range nodes {
		if _, exists := g[n]; !exists {
			return false, n
		}
	}

	return true, missingNode
}

func (g graphConcrete[T]) CountNodes() int {
	return len(g)
}

// If one of the nodes doesn't exist yet, it is created.
func (g graphConcrete[T]) AddEdge(node1, node2 T) {
	g.AddNode(node1)
	g.AddNode(node2)

	g.addEdgeOneDirection(node1, node2)
	g.addEdgeOneDirection(node2, node1)
}

func (g graphConcrete[T]) RemoveEdge(node1, node2 T) {
	if !g.HasEdge(node1, node2) {
		return
	}

	g[node1].Remove(node2)
	g[node2].Remove(node1)
}

// The caller is expected to ensure, that node1 and node2 exist.
func (g graphConcrete[T]) addEdgeOneDirection(node1, node2 T) {
	_, exists := g[node1]
	if !exists {
		g[node1] = NewSet[T]()
	}

	g[node1].Add(node2)
}

func (g graphConcrete[T]) HasEdge(node1, node2 T) bool {
	setRef, exists := g[node1]
	if !exists {
		return false
	}

	return setRef.Contains(node2)
}

func (g graphConcrete[T]) CountEdges() int {
	total := 0
	for _, neighbours := range g {
		total += neighbours.Count()
	}
	return total
}

func (g graphConcrete[T]) AdjacentNodes(node T) Set[T] {
	connectedNodes, _ := g[node]
	return connectedNodes
}

func (g graphConcrete[T]) AllNodes() Set[T] {
	outputSet := NewSet[T]()
	for node := range g {
		outputSet.Add(node)
	}
	return outputSet
}

func (g graphConcrete[T]) ConnectedNodes() Set[T] {
	outputSet := NewSet[T]()
	for node, neighbours := range g {
		if neighbours.Count() > 0 {
			outputSet.Add(node)
		}
	}
	return outputSet
}

func (g graphConcrete[T]) UnconnectedNodes() Set[T] {
	outputSet := NewSet[T]()
	for node, neighbours := range g {
		if neighbours.Count() == 0 {
			outputSet.Add(node)
		}
	}
	return outputSet
}
