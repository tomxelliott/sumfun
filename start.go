package main

import (
	"fmt"
	"bufio"
	"os"

	"quiz/game"
)

func main() {
	fmt.Println("Welcome! Try and answer as many sums in 60 seconds!")
	fmt.Println("When you are ready, press enter to begin the timer...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	game.Game()
}
