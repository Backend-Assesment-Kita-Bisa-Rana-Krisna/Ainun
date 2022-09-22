package main

import (
	"fmt"
)

type Bundle struct {
	Name  string
	Apple int
	Cake  int
}

func main() {
	bundle := Bundle{Name: "Ainun", Cake: 20, Apple: 25}
	boxCount := countBox(bundle.Cake, bundle.Apple)
	cakeCount, appleCount := countPartOfBox(bundle.Cake, bundle.Apple, boxCount)
	fmt.Printf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", bundle.Name, bundle.Cake, bundle.Apple)
	fmt.Printf("How many boxes that %v can make? %d box\n", bundle.Name, boxCount)
	fmt.Printf("how many cakes and apples every box have? %d cake and %d apple in every box\n", cakeCount, appleCount)
}

func countBox(cakes int, apples int) int {
	for i := cakes; i > 0; i-- {
		if cakes%i == 0 && apples%i == 0 {
			return i
		}
	}
	return 0
}

func countPartOfBox(cakes int, apples int, boxCount int) (cakeCount int, appleCount int) {
	cakeCount = cakes / boxCount
	appleCount = apples / boxCount
	return
}
