package graph

import (
	"github.com/JesseleDuran/osm-graph/coordinates"
	"github.com/JesseleDuran/osm-graph/json"
	"github.com/golang/geo/s2"
)

// Represents a set of nodes related between them.
type Graph struct {
	Nodes Nodes
}

// newEmptyGraph creates a graph with 0 vertices.
func newEmptyGraph() Graph {
	return Graph{Nodes: make(Nodes, 0)}
}

// BuildFromJsonFile creates a new graph from a json file.
func BuildFromJsonFile(path string, setWeight SetWeight) Graph {
	g := newEmptyGraph()
	decoder, file := json.NewDecoder(path)
	defer file.Close()
	for decoder.More() {
		nodes := make([]EncodedNode, 0)
		if decoder.Decode(&nodes) == nil {
			for i := 0; i < len(nodes)-1; i++ {
				g.RelateNodesByID(
					s2.CellID(nodes[i].CellId), s2.CellID(nodes[i+1].CellId), setWeight,
				)
			}
		}
	}
	return g
}

// Add note to the given graph.
func (g *Graph) AddNodes(nodes ...Node) {
	for _, n := range nodes {
		g.Nodes[n.ID] = n
	}
}

// FindOrCreateNode find a node on the graph by the given ID.
// if the node does not exists then is created.
func (g *Graph) FindNodeOrCreate(id s2.CellID) *Node {
	node, ok := g.Nodes[id]
	if !ok {
		node = Node{ID: id, Edges: make(map[s2.CellID]Edge)}
		g.AddNodes(node)
	}
	return &node
}

// RelateNodesByID relate two nodes using its IDs.
func (g *Graph) RelateNodesByID(a, b s2.CellID, setWeight SetWeight) {
	var w float64
	nodeA, nodeB := g.FindNodeOrCreate(a), g.FindNodeOrCreate(b)
	pointA := coordinates.FromS2LatLng(nodeA.ID.LatLng())
	pointB := coordinates.FromS2LatLng(nodeB.ID.LatLng())
	if setWeight == nil {
		w = coordinates.Distance(pointA, pointB)
	} else {
		w = setWeight(pointA, pointB)
	}

	// The relation of nodes is bi-directional.
	nodeA.Edges[nodeB.ID] = Edge{Weight: w}
	nodeB.Edges[nodeA.ID] = Edge{Weight: w}
}
