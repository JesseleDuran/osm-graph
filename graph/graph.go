package graph

import (
  "github.com/JesseleDuran/osm-graph/cell"
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
        g.RelateNodesByID(nodes[i], nodes[i+1], setWeight)
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
func (g *Graph) FindNodeOrCreate(e EncodedNode) *Node {
  node, ok := g.Nodes[s2.CellID(e.CellId)]
  if !ok {
    node = Node{ID: s2.CellID(e.CellId), Edges: make(map[s2.CellID]Edge), Name: e.Name}
    g.AddNodes(node)
  } else {
    if node.Name == "" && e.Name != "" {
      node.Name = e.Name
      g.Nodes[s2.CellID(e.CellId)] = node
    }
  }
  return &node
}

// RelateNodesByID relate two nodes using its IDs.
func (g *Graph) RelateNodesByID(a, b EncodedNode, setWeight SetWeight) {
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

func (g *Graph) FindNodeRecursive(cid s2.CellID) s2.CellID {
  if _, ok := g.Nodes[cid]; ok {
    return cid
  }
  neighbors := cell.Neighbors(cid, 1)
  for _, n := range neighbors {
    if _, ok := g.Nodes[n]; ok {
      return n
    }
  }
  return s2.CellID(0)
}

