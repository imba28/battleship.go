package battleship

import "testing"

func TestRound_AddPlayer(t *testing.T) {
	r := Round{}
	p1 := Player{}
	p2 := Player{}

	r.AddPlayer(&p1)
	if r.playerA != &p1 {
		t.Errorf("Adding a player to a round should set playerA.")
	}

	r.AddPlayer(&p2)
	if r.playerB != &p2 {
		t.Errorf("Adding a second player to a round should set playerB.")
	}

	r.AddPlayer(&p2)
	if r.playerA != &p1 {
		t.Errorf("Adding a player to a full round should not modify the struct.")
	}
	if r.playerB != &p2 {
		t.Errorf("Adding a player to a full round should not modify the struct.")
	}
}

func TestRound_Fail(t *testing.T) {
	t.Errorf("This test should fail!")
}

func TestRound_IsWaiting(t *testing.T) {
	r := Round{}
	p1 := Player{}
	p2 := Player{}

	if !r.IsWaiting() {
		t.Errorf("A round with no assgined players should be in wait state.")
	}

	r.AddPlayer(&p1)
	if !r.IsWaiting() {
		t.Errorf("A round with one assgined player should be in wait state.")
	}

	r.AddPlayer(&p2)
	if r.IsWaiting() {
		t.Errorf("A round with two assgined players should ready.")
	}
}
