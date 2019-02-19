package battleship

import (
	"testing"
)

func TestCoordinate(t *testing.T) {
	c := Coordinate{x: 5, y: 4}

	if c.x != 5 {
		t.Errorf("Wrong x coordinate, expected 5 but got %d", c.x)
	}
	if c.y != 4 {
		t.Errorf("Wrong y coordinate, expected 4 but got %d", c.y)
	}
}

func TestNewRandomCoordinate(t *testing.T) {
	maxValue := 10

	for i := 0; i < 15; i++ {
		c := NewRandomCoordinate(maxValue, maxValue)

		if c.x >= maxValue {
			t.Errorf("x coordinate does not respect bounds, expected the value to be less than %d but got %d", maxValue, c.x)
		}
		if c.y >= maxValue {
			t.Errorf("y coordinate does not respect bounds, expected the value to be less than %d but got %d", maxValue, c.y)
		}
	}
}

func TestCoordinate_Blocks(t *testing.T) {
	c := Coordinate{x: 1, y: 1}
	set := map[[2]Coordinate]bool{
		[2]Coordinate{
			{x: 5, y: 4},
			{x: 5, y: 4},
		}: true,
		[2]Coordinate{
			{x: 1, y: 2},
			{x: 2, y: 2},
		}: false,
		[2]Coordinate{
			{x: 5, y: 5},
			{x: 5, y: 4},
		}: false,
		[2]Coordinate{
			{x: 3, y: 3},
			{x: 3, y: 3},
		}: true,
		[2]Coordinate{
			{x: 0, y: 0},
			{x: 9, y: 9},
		}: false,
		[2]Coordinate{
			{x: 5, y: 5},
			{x: 5, y: 5},
		}: true,
		[2]Coordinate{
			c,
			c,
		}: true,
	}

	for item, blocks := range set {
		a := item[0]
		b := item[1]

		if a.Blocks(b) != blocks {
			not := "not "
			if blocks {
				not = ""
			}
			t.Errorf("(%d, %d) should %sblock (%d, %d)", a.x, a.y, not, b.x, b.y)
		}
	}
}

func TestCoordinate_InList(t *testing.T) {
	c := Coordinate{x: 1, y: 1}

	list := []Coordinate{
		{x: 8, y: 5},
		{x: 1, y: 1},
		{x: 0, y: 0},
		{x: 1, y: 2},
		{x: 5, y: 9},
		{x: 9, y: 9},
		{x: 3, y: 9},
		{x: 8, y: 7},
		{x: 4, y: 2},
		{x: 3, y: 3},
		{x: 2, y: 9},
		{x: 5, y: 4},
	}

	set := map[Coordinate]bool{
		Coordinate{x: 5, y: 4}: true,
		Coordinate{x: 0, y: 0}: true,
		Coordinate{x: 1, y: 2}: true,
		Coordinate{x: 5, y: 9}: true,
		Coordinate{x: 9, y: 9}: true,
		c:                      true,
		Coordinate{x: 1, y: 1}: true,
		Coordinate{x: 9, y: 1}: false,
		Coordinate{x: 6, y: 9}: false,
		Coordinate{x: 6, y: 5}: false,
		Coordinate{x: 4, y: 3}: false,
		Coordinate{x: 3, y: 4}: false,
	}

	for c, shouldInclude := range set {
		if c.InList(list) != shouldInclude {
			not := "not "
			if shouldInclude {
				not = ""
			}
			t.Errorf("(%d, %d) should %sbe included in list.", c.x, c.y, not)
		}
	}
}
