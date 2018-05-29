package game

import (
	"time"
	"fmt"
)

var score int

func Game() {
	timer := time.NewTimer(time.Minute)
	counter := 1
	next := make(chan int)

	go func() {
		for {
			score += Start(counter)
			counter++
		}
		next <- 1
	}()

	select {
	case <- next:
	case <-timer.C:
		fmt.Println()
		fmt.Printf("Your final score was %d.\n", score)
		fmt.Println("Thanks for playing!")
	}
}
