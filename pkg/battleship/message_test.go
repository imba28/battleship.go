package battleship

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewAnnouncement(t *testing.T) {
	m := NewAnnouncement("foobar")
	if m.Name != MESSAGE_INFO {
		t.Errorf("Message type should be %s but is %s", MESSAGE_INFO, m.Name)
	}
	if m.Body != "foobar" {
		t.Errorf("Message content should be %s but is %s", "foobar", m.Body)
	}
}

func TestNewDrawBoard(t *testing.T) {
	ships := []Ship{
		{shipType: SHIP_CARRIER, coordinates: []Coordinate{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		}},
		{shipType: SHIP_DESTROYER, coordinates: []Coordinate{
			{5, 5},
			{5, 4},
			{5, 3},
		}},
		{shipType: SHIP_DESTROYER, coordinates: []Coordinate{
			{9, 9},
			{9, 8},
			{9, 7},
		}},
	}
	m := NewDrawBoard(ships)
	expected := "C00102030D555453D999897"

	if s := m.String(); !strings.Contains(s, fmt.Sprintf(`"%s"`, expected)) {
		t.Errorf("Expected message body to contain %s but got %s", expected, s)
	}
}

func TestMessage_String(t *testing.T) {
	m := NewAnnouncement("foobar")
	s := m.String()

	expected := fmt.Sprintf(`{"Name":"%s","Body":"%s","Time":%d}`, MESSAGE_INFO, m.Body, m.Time)
	if s != expected {
		t.Errorf("Expected marshaled message to be %s but got %s", expected, s)
	}

	m = NewAnnouncement("Row, row, row your boat gently down the stream.")
	s = m.String()

	expected = fmt.Sprintf(`{"Name":"%s","Body":"%s","Time":%d}`, MESSAGE_INFO, m.Body, m.Time)
	if s != expected {
		t.Errorf("Expected marshaled message to be %s but got %s", expected, s)
	}
}
