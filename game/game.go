package game

import (
	"math/rand"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"time"
	"reflect"
)

func Start(q int) int {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Question %d: ", q)
	answer := Question()

	scanner.Scan()
	input := scanner.Text()
	attempt, _ := strconv.Atoi(input)
	val := reflect.TypeOf(attempt)
	if k := val.Kind(); k != reflect.Int {
		return 0
	}

	return Evaluate(attempt, answer)
}

func Question() int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)

	fmt.Printf("%d + %d\n", x, y)
	return x + y
}

func Evaluate(attempt, answer int) int {
	if attempt == answer {
		return 1
	} else {
		return 0
	}
}
