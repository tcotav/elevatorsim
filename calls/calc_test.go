package calls 

import (
	"testing"
)

func TestAddStop(t *testing.T) {
	c := NewCallList(1, 10)
	c.AddStop(3)
	c.AddStop(5)
	c.AddStop(1)
	if c.UpList[0] != 3 {
		t.Errorf("UpList[0] = %d; want 3", c.UpList[0])
	}
	if c.UpList[1] != 5 {
		t.Errorf("UpList[1] = %d; want 5", c.UpList[1])
	}
}

func TestGetStops(t *testing.T) {
	c := NewCallList(1,10)
	c.AddStop(3)
	c.AddStop(5)
	stops := c.GetStops()
	if stops[0] != 3 {
		t.Errorf("stops[0] = %d; want 3", stops[0])
	}
	if stops[1] != 5 {
		t.Errorf("stops[1] = %d; want 5", stops[1])
	}
}

func TestMove(t *testing.T) {
	c := NewCallList(1,10)
	c.AddStop(3)
	c.AddStop(5)
	c.Move()
	if c.CurrentFloor != 2 {
		t.Errorf("CurrentFloor = %d; want 3", c.CurrentFloor)
		t.Error(c.GetStops(), c.Direction, c.CurrentFloor)
	}
	c.AddStop(1) // opposite direction
	c.Move()
	c.Move()
	if c.CurrentFloor != 4 {
		t.Errorf("CurrentFloor = %d; want 4", c.CurrentFloor)
		t.Error(c.GetStops(), c.Direction, c.CurrentFloor)
	}
	c.Move()
	if c.CurrentFloor != 5 {
		t.Errorf("CurrentFloor = %d; want 5", c.CurrentFloor)
		t.Error(c.GetStops(), c.Direction, c.CurrentFloor)
	}
	c.Move()
	if c.Direction != -1 {
		t.Errorf("Direction = %d; want -1", c.Direction)
		t.Error(c.GetStops(), c.Direction, c.CurrentFloor)
	}

	if c.CurrentFloor != 4{
		t.Errorf("CurrentFloor = %d; want 4", c.CurrentFloor)
		t.Error(c.GetStops(), c.Direction, c.CurrentFloor)
	}

}