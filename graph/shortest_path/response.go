package shortest_path

import "github.com/JesseleDuran/osm-graph/coordinates"

type Response struct {
  Steps       Steps
  TotalWeight float64 //might be distance
  Polyline    [][2]float64
}

type Step struct {
  Weight        float64
  StartAddress  string
  EndAddress    string
  StartLocation coordinates.Coordinates
  EndLocation   coordinates.Coordinates
}

type Steps []Step
