package bfs

import (
	"log"
	"testing"
	"time"

	"github.com/JesseleDuran/osm-graph/graph"
	"github.com/golang/geo/s2"
)

func TestBFS_Path(t *testing.T) {

	start := time.Now()
	g := graph.BuildFromJsonFile("testdata/osm-Graph-sp-16.json", nil)
	end := time.Since(start)
	log.Println("done Graph", end.Milliseconds(), len(g.Nodes))
	g.Nodes.ToGeoJSON()
	time.Sleep(time.Hour)

	bfs := BFS{Graph: g}
	//s := s2.CellIDFromToken("94ce50b26c")
	s := s2.CellIDFromToken("94ce50a5b")
	log.Println(s.LatLng().Lat.Degrees(), s.LatLng().Lng.Degrees())
	start1 := time.Now()
	bfs.Path(s, 10000)
	end1 := time.Since(start1)
	log.Println("done bfs", end1.Milliseconds(), len(g.Nodes))
	time.Sleep(time.Hour)
}
