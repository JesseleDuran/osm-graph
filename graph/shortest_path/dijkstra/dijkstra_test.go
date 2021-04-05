package dijkstra

import (
  "fmt"
  "log"
  "testing"
  "time"

  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/golang/geo/s2"
)

func TestDijkstra(t *testing.T) {
  start := time.Now()
  g := graph.BuildFromJsonFile("testdata/osm-graph-sp-16.json", nil)
  end := time.Since(start)
  log.Println("done graph", end.Milliseconds(), len(g.Nodes))
  //for n, _ := range g.Nodes {
  //log.Println("node", n.ToToken())
  //}

  s := s2.CellIDFromToken("94ce4fe4d")
  e := s2.CellIDFromToken("94ce5bb29")
  d := Dijkstra{Graph: g}
  _, prev := d.FromCellIDs(s, e)
  log.Println(len(prev))
  path := path(s, e, prev)
  for _, v := range path {
    fmt.Println(
      s2.CellFromCellID(v).ID().ToToken())
  }
}

func TestPath(t *testing.T) {
  previous := map[s2.CellID]s2.CellID{
    1: 0,
    2: 1,
    3: 2,
    4: 3,
    5: 4,
  }
  path := path(1, 5, previous)
  log.Println(path)
}
