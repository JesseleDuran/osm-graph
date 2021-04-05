package bfs

import (
  "container/list"

  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/golang/geo/s2"
)

type BFS struct {
  Graph graph.Graph
}

func (g BFS) PathRadius(start s2.CellID, m float64) graph.Nodes {
  result := make(graph.Nodes, 0)
  visited := make(map[s2.CellID]float64)
  queue := list.New()
  queue.PushBack(start)
  visited[start] = 0

  for queue.Len() > 0 {
    qnode := queue.Front()
    queue.Remove(qnode)
    cellID := qnode.Value.(s2.CellID)
    for k, e := range g.Graph.Nodes[cellID].Edges {
      if _, ok := visited[k]; ok {
        continue
      }
      currentWeight := visited[cellID] + e.Weight
      if currentWeight < m {
        visited[k] += currentWeight
        queue.PushBack(k)
        result[k] = g.Graph.Nodes[k]
      }
    }
  }
  return result
}

func (g BFS) Path(start, end s2.CellID) graph.Nodes {
  result := make(graph.Nodes, 0)
  visited := make(map[s2.CellID]bool)
  queue := list.New()
  queue.PushBack(start)
  visited[start] = true

  for queue.Len() > 0 {
    qnode := queue.Front()
    queue.Remove(qnode)
    cellID := qnode.Value.(s2.CellID)
    if cellID.ToToken() == end.ToToken() {
      return result
    }
    for k, _ := range g.Graph.Nodes[cellID].Edges {
      if _, ok := visited[k]; ok {
        continue
      }
      visited[k] = true
      queue.PushBack(k)
      result[k] = g.Graph.Nodes[k]
    }
  }
  return result
}
