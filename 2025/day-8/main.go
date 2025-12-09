package main

import (
	"fmt"
	"github.com/tienanr/advent-of-code/utils"
	"sort"
	"strings"
)

func parseLine(line string) []int {
	fields := strings.Split(line, ",")
	point := []int{}
	for _, f := range fields {
		point = append(point, utils.ParseInt(f))
	}
	return point
}

type Edge struct {
	fst int
	snd int
}

func solve(fn string, k int) {
	// read input
	points := [][]int{}
	for line := range utils.ReadFile(fn) {
		points = append(points, parseLine(line))
	}

	// list all edges
	edges := []Edge{}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{i, j})
		}
	}

	// for k closest edges
	distance_square := func(p Edge) int {
		dist := 0
		for i := 0; i < 3; i++ {
			d := points[p.fst][i] - points[p.snd][i]
			dist += d * d
		}
		return dist
	}
	sort.Slice(edges, func(i, j int) bool {
		return distance_square(edges[i]) < distance_square(edges[j])
	})

	solvePart1(edges, k)
	solvePart2(edges, points)
}

func solvePart1(edges []Edge, k int) {
	// connect k shortest edges
	edges = edges[:k]

	// construct Graph
	g := NewGraph()
	for _, edge := range edges {
		g.AddEdge(edge.fst, edge.snd)
		g.AddEdge(edge.snd, edge.fst)
	}

	// find all components
	components := g.GetConnectedComponents()

	// find size of 3 larges components
	sizes := []int{}
	for _, c := range components {
		sizes = append(sizes, len(c))
	}
	sort.Ints(sizes)
	sizes = sizes[len(sizes)-3:len(sizes)]

	// print answer
	ans1 := 1
	for _, size := range sizes {
		ans1 *= size
	}
	fmt.Println("part 1 solution:", ans1)
}

func solvePart2(edges []Edge, points [][]int) {
	d := NewDSU(len(points))
	var ans2 int

	for _, edge := range edges {
		d.Union(edge.fst, edge.snd)
		if d.size == 1 {
			ans2 = points[edge.fst][0] * points[edge.snd][0]
			break
		}
	}
	fmt.Println("part 2 solution:", ans2)
}

func main() {
	solve("example.txt", 10)
	//solve("question.txt", 1000)
}

// Graph implementation
type Graph struct {
	Adj map[int][]int
}

func NewGraph() Graph {
	return Graph{make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	_, ok := g.Adj[u]
	if !ok {
		g.Adj[u] = []int{}
	}
	g.Adj[u] = append(g.Adj[u], v)
}

func (g Graph) GetConnectedComponents() [][]int {
	visited := make(map[int]bool)
	components := [][]int{}

	for v := range g.Adj {
		if !visited[v] {
			component := []int{}
			dfs(g, v, visited, &component)
			components = append(components, component)
		}
	}

	return components
}

func dfs(g Graph, v int, visited map[int]bool, component *[]int) {
	visited[v] = true
	*component = append(*component, v)
	for _, u := range g.Adj[v] {
		if !visited[u] {
			dfs(g, u, visited, component)
		}
	}
}

// Disjoint Set Union implementation
type DSU struct {
	parent []int
	rank []int
	size int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		rank: make([]int, n),
		size: n,
	}
	for i := 0; i<n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 0
	}
	return dsu
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i]) // compression
	return d.parent[i]
}

func (d *DSU) Union(i, j int) {
	root_i := d.Find(i)
	root_j := d.Find(j)

	if root_i != root_j { // merge 2 sets
		if d.rank[root_i] < d.rank[root_j] {
			d.parent[root_i] = root_j
		} else if d.rank[root_i] > d.rank[root_j] {
			d.parent[root_j] = root_i
		} else {
			d.parent[root_i] = root_j
			d.rank[root_j]++
		}
		d.size--
	}
}
