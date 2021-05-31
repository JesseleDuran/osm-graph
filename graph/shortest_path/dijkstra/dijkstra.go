package dijkstra

import (
  "math"

  "github.com/JesseleDuran/osm-graph/coordinates"
  "github.com/JesseleDuran/osm-graph/graph"
  "github.com/JesseleDuran/osm-graph/graph/shortest_path"
  "github.com/JesseleDuran/osm-graph/graph/shortest_path/dijkstra/heap"
  "github.com/golang/geo/s2"
)

const INFINITE = math.MaxInt64

type Dijkstra struct {
  Graph graph.Graph
}

type Previous map[s2.CellID]s2.CellID
type PathWeight map[s2.CellID]float64

func (d Dijkstra) FromCoordinates(origin, destiny coordinates.Coordinates) shortest_path.Response {
  originCell := d.Graph.FindNodeRecursive(origin.ToCellID())
  destinyCell := d.Graph.FindNodeRecursive(destiny.ToCellID())
  weight, prev := d.FromCellIDs(originCell, destinyCell)

  return shortest_path.Response{
    Steps:       Steps(path(originCell, destinyCell, prev), d.Graph),
    TotalWeight: weight[destinyCell],
    Polyline:    pathPolyline(originCell, destinyCell, prev),
  }
}

func (d Dijkstra) FromCellIDs(start, end s2.CellID) (PathWeight, Previous) {
  //maps from each node to the total weight of the total shortest path.
  pathWeight := make(PathWeight, 0)

  //maps from each node to the previous node in the "current" shortest path.
  previous := make(Previous, 0)

  remaining := heap.Create()
  // insert first node id the PQ, the start node.
  remaining.Insert(heap.Node{Value: start, Cost: 0})

  // initialize pathWeight all to infinite value.
  for _, v := range d.Graph.Nodes {
    pathWeight[v.ID] = INFINITE
  }
  //start node distance to itself is 0.
  pathWeight[start] = 0

  //the previous node does not exists
  previous[start] = INFINITE

  visit := make(map[s2.CellID]bool, 0)

  //while the PQ is not empty.
  for !remaining.IsEmpty() {
    // extract the min value of the PQ.
    min, _ := remaining.Min()
    visit[min.Value] = true
    remaining.DeleteMin()
    if min.Value.ToToken() == end.ToToken() {
      return pathWeight, previous
    }
    // if the node has edges, the loop through it.
    if v, ok := d.Graph.Nodes[min.Value]; ok {
      //change to normal for
      for nodeNeighbor, e := range v.Edges {

        if visit[nodeNeighbor] {
          continue //change to negative condition
        }
        visit[nodeNeighbor] = true

        // the value is the one of the current node plus the weight(a, neighbor)
        currentPathValue := pathWeight[min.Value] + e.Weight

        if currentPathValue < pathWeight[nodeNeighbor] {
          pathWeight[nodeNeighbor] = currentPathValue
          previous[nodeNeighbor] = min.Value
        }
        remaining.Insert(heap.Node{Value: nodeNeighbor, Cost: currentPathValue})
      }
    }
  }
  return pathWeight, previous
}

//key : end, value: prev
func path(start, end s2.CellID, previous Previous) []s2.CellID {
  result := make([]s2.CellID, 0)
  result = append(result, end)
  var prev s2.CellID
  _, startOk := previous[start]
  _, endOk := previous[end]
  if !startOk && !endOk {
    return result
  }

  for prev != start {
    prev = previous[end]
    result = append(result, prev)
    end = prev
    //log.Println(prev, end)
  }

  resultSorted := make([]s2.CellID, len(result))
  j := 0
  for i := len(result) - 1; i >= 0; i-- {
    resultSorted[j] = result[i]
    j++
  }
  return resultSorted
}

//key : end, value: prev
func pathPolyline(start, end s2.CellID, previous Previous) [][2]float64 {
  result := make([][2]float64, 0)
  result = append(result, [2]float64{
    end.LatLng().Lng.Degrees(),
    end.LatLng().Lat.Degrees(),
  })
  var prev s2.CellID
  _, startOk := previous[start]
  _, endOk := previous[end]
  if !startOk && !endOk {
    return result
  }

  for prev != start {
    prev = previous[end]
    result = append(result, [2]float64{
      prev.LatLng().Lng.Degrees(),
      prev.LatLng().Lat.Degrees(),
    })
    end = prev
    //log.Println(prev, end)
  }

  resultSorted := make([][2]float64, len(result))
  j := 0
  for i := len(result) - 1; i >= 0; i-- {
    resultSorted[j] = result[i]
    j++
  }
  return resultSorted
}

func Steps(path []s2.CellID, graph graph.Graph) shortest_path.Steps {
  result := make(shortest_path.Steps, 0)
  for i := 0; i < len(path)-1; i++ {
    start := graph.Nodes[path[i]]
    end := graph.Nodes[path[i+1]]
    result = append(result, shortest_path.Step{
      Weight:       0,
      StartAddress: start.Name,
      EndAddress:   end.Name,
      StartLocation: coordinates.Coordinates{
        Lat: start.ID.LatLng().Lat.Degrees(),
        Lng: start.ID.LatLng().Lng.Degrees(),
      },
      EndLocation: coordinates.Coordinates{
        Lat: end.ID.LatLng().Lat.Degrees(),
        Lng: end.ID.LatLng().Lng.Degrees(),
      },
    })
  }

  return result
}
