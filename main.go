package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Mari")
	sayGreeting("Niha")
	sayBye("mari")

	cycleNames([]string{"noman", "farret", "tifa"}, sayGreeting)
	cycleNames([]string{"noman", "farret", "tifa"}, sayBye)

	a1 := circleArea(30.5)
	a2 := circleArea(25)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}
