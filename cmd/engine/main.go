package main

import (
	"time"
	"math/rand"
	"github.com/tcotav/elevsim/calls"
	"fmt"
)

const MOVE_SLEEP = 2 * time.Second
const EVENT_CHANCE_PERCENT = 80

func mainloop(c *calls.CallList) {
	for {
		c.Move()
		time.Sleep(MOVE_SLEEP)
	}
}

func randomCall(c *calls.CallList) {
	eventChance := rand.Intn(100) 
	if eventChance <= EVENT_CHANCE_PERCENT {
		// create a random call
		floor := rand.Intn(c.MaxFloor-c.MinFloor) + c.MinFloor 
		c.AddStop(floor)
	}
}

func main() {
	minFloor := 1
	maxFloor := 10
	fmt.Printf("Starting on floor %d\n", minFloor)
	c := calls.NewCallList(minFloor, maxFloor)
	go mainloop(c)
	// create some random elevator calls
	for {
		randomCall(c)
		time.Sleep(2 * MOVE_SLEEP)
	}
}