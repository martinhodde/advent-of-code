package solutions

import (
	"aoc-2025/internal/registry"
	"container/heap"
	"fmt"
	"slices"
)

// Number of permitted junction connections
const maxConnections = 1000

// Connection represents a weighted edge between two junctions in 3D space.
type Connection struct {
	from     [3]int
	to       [3]int
	distance int
}

// PriorityQueue implements a min-heap for connections between 3D positions based on distance.
type PriorityQueue []Connection

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Connection))
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// UnionFind implements the union-find (disjoint set) data structure on 3D positions.
type UnionFind struct {
	parent map[[3]int][3]int
}

func NewUnionFind(positions [][3]int) *UnionFind {
	uf := &UnionFind{parent: make(map[[3]int][3]int)}
	for _, pos := range positions {
		uf.parent[pos] = pos
	}
	return uf
}
func (uf *UnionFind) Find(x [3]int) [3]int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *UnionFind) Union(a, b [3]int) {
	rootA := uf.Find(a)
	rootB := uf.Find(b)
	if rootA != rootB {
		uf.parent[rootB] = rootA
	}
}

func init() {
	registry.Register(8, 1, SolveDay8Part1)
	registry.Register(8, 2, SolveDay8Part2)
}

func SolveDay8Part1(input []string) string {
	positions := ParseJunctionPositions(input)
	connections := MakeConnections(positions)
	numCircuits := 3
	largestCircuits := GetKLargestCircuits(connections, numCircuits)
	product := CircuitSizeProduct(largestCircuits)
	return fmt.Sprintf("The product of the sizes of the %d largest circuits is %d", numCircuits, product)
}

func SolveDay8Part2(input []string) string {
	return ""
}

// CircuitSizeProduct computes the product of the sizes of the provided circuits.
func CircuitSizeProduct(circuits [][][3]int) int {
	product := 1
	for _, circuit := range circuits {
		product *= len(circuit)
	}

	return product
}

// GetKLargestCircuits identifies the k largest circuits (connected components)
// in the junction graph represented as an adjacency list.
func GetKLargestCircuits(components map[[3]int][][3]int, k int) [][][3]int {
	var circuits [][][3]int
	for _, nodes := range components {
		circuits = append(circuits, nodes)
	}

	// Sort descending by size
	slices.SortFunc(circuits, func(a, b [][3]int) int {
		return len(b) - len(a)
	})

	// Trim to k largest if necessary
	if k < len(circuits) {
		return circuits[:k]
	}
	return circuits
}

// MakeConnections constructs a graph of junction connections using Kruskal's minimum spanning
// tree algorithm, subject to the global limit on the number of connections.
func MakeConnections(positions [][3]int) map[[3]int][][3]int {
	// Build a min-heap of all possible pairwise connections between junctions by distance
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i, posA := range positions {
		for j := i + 1; j < len(positions); j++ {
			posB := positions[j]
			dx, dy, dz := posA[0]-posB[0], posA[1]-posB[1], posA[2]-posB[2]
			distSquared := dx*dx + dy*dy + dz*dz // Use squared distance to avoid float operations
			heap.Push(pq, Connection{posA, posB, distSquared})
		}
	}

	// Connect junctions using union-find until reaching the max allowed connections
	uf := NewUnionFind(positions)
	numConnections := 0
	for pq.Len() > 0 && numConnections < maxConnections {
		conn := heap.Pop(pq).(Connection)
		uf.Union(conn.from, conn.to)
		numConnections++
	}

	// Build component graph as adjacency list
	components := make(map[[3]int][][3]int)
	for _, pos := range positions {
		root := uf.Find(pos)
		components[root] = append(components[root], pos)
	}

	return components
}

// ParseJunctionPositions parses a list of strings representing 3D coordinates
// into a slice of integer triplets.
func ParseJunctionPositions(input []string) [][3]int {
	var positions [][3]int
	for _, line := range input {
		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err != nil {
			continue // Skip invalid lines (none present in input)
		}
		positions = append(positions, [3]int{x, y, z})
	}

	return positions
}
