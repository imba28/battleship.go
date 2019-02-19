package battleship

import "testing"

func TestPlayer_AddShips(t *testing.T) {
	p := Player{}
	p.AddShips()

	if len(p.ships) <= 0 {
		t.Fatal("AddShips should add ships to player struct!")
	}

	shipTypes := [...]rune{SHIP_CARRIER, SHIP_BATTLESHIP, SHIP_DESTROYER, SHIP_SUBMARINE}

	for i, s := range p.ships {
		if shipTypes[i] != s.shipType {
			t.Errorf("Ship number %d should be of type %c but is %c", i, shipTypes[i], s.shipType)
		}

		for j, ss := range p.ships {
			if i == j {
				continue
			}

			for _, coord := range s.coordinates {
				if coord.InList(ss.coordinates) {
					t.Errorf("%v collides with another ship %v!", coord, ss.coordinates)
				}
			}
		}
	}
}
