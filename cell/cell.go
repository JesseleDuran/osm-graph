package cell

import "github.com/golang/geo/s2"

var left = func(cid s2.CellID) s2.CellID { return cid.EdgeNeighbors()[0] }
var down = func(cid s2.CellID) s2.CellID { return cid.EdgeNeighbors()[1] }
var right = func(cid s2.CellID) s2.CellID { return cid.EdgeNeighbors()[2] }
var up = func(cid s2.CellID) s2.CellID { return cid.EdgeNeighbors()[3] }
// Vector transformations to apply to a cell.
var transformations = [4]func(cid s2.CellID) s2.CellID{right, down, left, up}

// Neighbors Retrieves the neighbors of the given cell at a given level.
// The eight rectangles around the cell 'X' represents neighbors
// of level 1:
//
// +---+---+---+
// |   |   |   |
// +-----------+
// |   | X |   |
// +-----------+
// |   |   |   |
// +---+---+---+
//
func Neighbors(cid s2.CellID, lvl int) []s2.CellID {
  i, limit, t := 1, (2*lvl)+1, 0
  result := make([]s2.CellID, 0)
  for i := 0; i < lvl; i++ {
    cid = left(cid)
    cid = up(cid)
  }
  for k := 0; k < 4*(limit-1); k++ {
    if i%limit == 0 {
      t++
      i = 1
    }
    i++
    cid = transformations[t](cid)
    result = append(result, cid)
  }
  return result
}
