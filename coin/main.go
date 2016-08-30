package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var events float64
	var userChoice int
	flag.Float64Var(&events, "events", 1000.00, "events number to test")
	flag.IntVar(&userChoice, "choice", 0, "choice 0 or 1")
	flag.Parse()

	winsCases := 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < int(events); i++ {
		if rand.Intn(2) == userChoice {
			winsCases++
		}
	}
	if winsCases >= (int(events) / 2) {
		fmt.Println("You win on ", float64(float64(winsCases*100)/events), "% cases")
	} else {
		fmt.Println("You lose on ", +float64(float64(winsCases*100)/events), "% cases")
	}
}
