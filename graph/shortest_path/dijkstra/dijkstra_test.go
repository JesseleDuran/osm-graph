package dijkstra

import (
	"log"
	"testing"
	"time"

	"github.com/JesseleDuran/osm-graph/graph"
	"github.com/golang/geo/s2"
)

func TestDijkstra(t *testing.T) {
	start := time.Now()
	g := graph.BuildFromJsonFile("testdata/osm-graph-sp.json")
	end := time.Since(start)
	log.Println("done graph", end.Milliseconds(), len(g.Nodes))
	g.Nodes.ToGeoJSON()
	time.Sleep(time.Hour)
	//d := Dijkstra{graph: g}
	//s := s2.CellIDFromToken("94ce595164")
	//e := s2.CellIDFromToken("94ce50b26c")
	//_, prev := d.FromCellIDs(s, e)
	//log.Println(prev)
	//for k, v := range prev {
	//  fmt.Println(s2.CellFromCellID(k).ID().ToToken(),
	//    s2.CellFromCellID(v).ID().ToToken())
	//}

	//fmt.Println("----------------------------")
	//p := path(s, e, prev)
	//fmt.Println("ola", p)
	//fmt.Println("----------------------------")

	//log.Println(Path(7126102609, r[0], prev))
	//c := g.Nodes[r[0]]

	//cc := CoordinatesPath(coordinates.Coordinates{
	//  Lat: 6.2451111,
	//  Lng: -75.5907088,
	//}, c.Point, prev, g)
	//for i := range cc {
	//  fmt.Println("[", cc[i].Lng, ",", cc[i].Lat, "],")
	//}
	//log.Println(weight[r[0]])

	//for i := 0; i <= 100; i++ {
	//  weight, prev := Dijkstra(7126102609, 0, g)
	//  fmt.Println("----------------------------")
	//  r := make([]int, 0)
	//  for k, v := range weight {
	//    if v != math.MaxInt64 {
	//      //fmt.Println("k", k, v)
	//      //fmt.Println("rev", prev[k])
	//      r = append(r, prev[k])
	//    }
	//  }
	//  sort.Ints(r)
	//  fmt.Println(r)
	//  fmt.Println("----------------------------")
	//}
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
