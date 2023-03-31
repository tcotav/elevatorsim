package calls

import (
	"fmt"
	"sort"
	"sync"
)

/*
Designing an elevator car call control system

Specifically, the algo for finding what is the next floor to go to.
- if there are no calls, go to the ground floor
- if there are calls, go to the closest call in the same direction already in progress
- else go to the closest call in the opposite direction and change direction
*/

// when direction change, we reverse the order of the calls

type CallList struct {
	mx		   sync.Mutex
	UpList       []int
	DownList     []int
	Direction    int
	CurrentFloor int
	MaxFloor	 int
	MinFloor 	int
}

func NewCallList(minFloor int, maxFloor int) *CallList {
	return &CallList{
		Direction:    0,
		CurrentFloor: minFloor,
		MaxFloor:	 maxFloor,
		MinFloor: minFloor,
		UpList:       make([]int, 0),
		DownList:     make([]int, 0),
	}
}

func isDupeCall(x int, list []int) bool {
	for _, v := range list {
		if v == x {
			return true
		}
	}
	return false
}

func (c *CallList) AddStop(floor int) {
	if floor > c.MaxFloor || floor < c.MinFloor {
		fmt.Println("invalid floor")
		return
	}
	if floor == c.CurrentFloor {
		return
	}
	c.mx.Lock()
	defer c.mx.Unlock()
	if floor > c.CurrentFloor {
		if c.Direction == 0 {
			c.Direction = 1
		}
		if !isDupeCall(floor, c.UpList) {
			fmt.Printf("adding floor %d\n", floor)
			c.UpList = append(c.UpList, floor)
			sort.Ints(c.UpList)
		}
	} else {
		if c.Direction == 0 {
			c.Direction = -1
		}
		if !isDupeCall(floor, c.DownList) {
			fmt.Printf("adding floor %d\n", floor)
			c.DownList = append(c.DownList, floor)
			// reverse order sort
			sort.Sort(sort.Reverse(sort.IntSlice(c.DownList)))
		}
	}
}

func (c *CallList) GetStops() []int {
	// order is dependent on direction of travel
	if c.Direction == 1 {
		return append(c.UpList, c.DownList...)
	} else {
		return append(c.DownList, c.UpList...)
	}
}

func (c *CallList) Move() {
	c.mx.Lock()
	defer c.mx.Unlock()
	if len(c.UpList) == 0 && len(c.DownList) == 0 {
		c.Direction = 0
		fmt.Println("no more stops")
		return
	}

	// move in the same direction as we're already going until we've emptied that list
	if c.Direction == 1 {
		if len(c.UpList) > 0 {
			// keep moving towards next stop
			c.CurrentFloor += 1
			fmt.Println("moving up to floor", c.CurrentFloor)
			if c.CurrentFloor == c.UpList[0] {
				fmt.Println("Stopping at floor", c.CurrentFloor)
				c.UpList = c.UpList[1:]
			}
		} else {
			c.Direction = -1
		}
	} 
	if c.Direction == -1 {
		if len(c.DownList) > 0 {
			c.CurrentFloor -= 1
			fmt.Println("moving down to floor", c.CurrentFloor)
			if c.CurrentFloor == c.DownList[0] {
				fmt.Println("Stopping at floor", c.CurrentFloor)
				c.DownList = c.DownList[1:]
				if len(c.DownList) == 0 {
					c.Direction = 0
				}
			}
		} else {
			c.Direction = 1
		}
	}
	if c.CurrentFloor > c.MaxFloor {
		c.CurrentFloor = c.MaxFloor
		c.Direction = -1
	} else if c.CurrentFloor < c.MinFloor {
		c.CurrentFloor = c.MinFloor
		c.Direction = 1
	}

	fmt.Println(c.GetStops())
}


