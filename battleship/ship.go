package battleship

import "math/rand"

const (
	SHIP_CARRIER           = 'C'
	SHIP_BATTLESHIP        = 'B'
	SHIP_DESTROYER         = 'D'
	SHIP_SUBMARINE         = 'S'
	ORIENTATION_VERTICAL   = 'v'
	ORIENTATION_HORIZONTAL = 'h'
)

type Ship struct {
	shipType    rune
	coordinates []Coordinate
	hits        []Coordinate
}

func NewRandomShip(shipType rune, blockedTiles *[]Coordinate) Ship {
	ship := Ship{shipType: shipType}
	var l int

	switch shipType {
	case SHIP_CARRIER:
		l = 5
	case SHIP_BATTLESHIP:
		l = 4
	case SHIP_DESTROYER:
		l = 3
	case SHIP_SUBMARINE:
		l = 2
	}

	for {
		orientation := ORIENTATION_VERTICAL
		if rand.Intn(2) == 1 {
			orientation = ORIENTATION_HORIZONTAL
		}

		var maxX, maxY int

		switch orientation {
		case ORIENTATION_HORIZONTAL:
			maxX = BOARD_SIZE - l
			maxY = BOARD_SIZE
		case ORIENTATION_VERTICAL:
			maxX = BOARD_SIZE
			maxY = BOARD_SIZE - l
		}

		headCoordinate := NewRandomCoordinate(maxX, maxY)
		if headCoordinate.InList(*blockedTiles) {
			continue
		}

		coords := make([]Coordinate, l)
		coords[0] = headCoordinate

		for i := 1; i <= l-1; i++ {
			switch orientation {
			case ORIENTATION_HORIZONTAL:
				coords[i] = Coordinate{x: headCoordinate.x + i, y: headCoordinate.y}
			case ORIENTATION_VERTICAL:
				coords[i] = Coordinate{x: headCoordinate.x, y: headCoordinate.y + i}
			}
		}

		shipBlocksExistingShip := false
		for _, coord := range coords {
			if coord.InList(*blockedTiles) {
				shipBlocksExistingShip = true
				break
			}
		}

		if shipBlocksExistingShip {
			continue
		}

		ship.coordinates = coords
		break
	}

	return ship
}
