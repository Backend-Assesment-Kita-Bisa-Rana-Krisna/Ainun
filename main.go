package main

import (
	"fmt"
)

var (
	name  string
	apple int
	cake  int
)

func main() {
	name = "Ainun"
	cake = 20
	apple = 25
	boxCount := countBox(cake, apple)
	cakeCount, appleCount := countPartOfBox(cake, apple, boxCount)
	response := Response(name, cake, apple, boxCount, cakeCount, appleCount)
	fmt.Println(response)
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

func Response(name string, cake int, apple int, boxCount int, cakeCount int, appleCount int) string{
	var response string
	response += fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", name, cake, apple)
	response += fmt.Sprintf("How many boxes that %v can make? %d box\n", name, boxCount)
	response += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	return response
}