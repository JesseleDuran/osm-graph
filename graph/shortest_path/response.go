package shortest_path

import "github.com/JesseleDuran/osm-graph/coordinates"

type Response struct {
  Steps       Steps        `json:"steps"`
  TotalWeight float64      //might be distance
  Polyline    [][2]float64 `json:"polyline"`
}

type Step struct {
  Weight        float64
  StartAddress  string                  `json:"start_address"`
  EndAddress    string                  `json:"end_address"`
  StartLocation coordinates.Coordinates `json:"start_location"`
  EndLocation   coordinates.Coordinates `json:"end_location"`
}

type Steps []Step
