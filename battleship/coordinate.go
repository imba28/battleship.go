package battleship

import "math/rand"

type Coordinate struct {
	x, y int
}

type CoordinateList []Coordinate

func (c *Coordinate) Blocks(other Coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func (c *Coordinate) InList(list []Coordinate) bool {
	for _, other := range list {
		if c.Blocks(other) {
			return true
		}
	}

	return false
}

func NewRandomCoordinate(maxX int, maxY int) Coordinate {
	x := rand.Intn(maxX)
	y := rand.Intn(maxY)

	return Coordinate{x, y}
}
