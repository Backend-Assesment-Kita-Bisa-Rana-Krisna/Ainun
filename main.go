package main

import (
	"fmt"
)

type Bundle struct {
	Name  string
	Apple int
	Cake  int
}

func (u *Bundle) CountBox() int {
	for i := u.Cake; i > 0; i-- {
		if u.Cake%i == 0 && u.Apple%i == 0 {
			return i
		}
	}
	return 0
}

func (u *Bundle) CountPartOfBox() (cakeCount int, appleCount int) {
	cakeCount = u.Cake / u.CountBox()
	appleCount = u.Apple / u.CountBox()
	return
}

func main() {
	bundle := Bundle{Name: "Ainun", Cake: 20, Apple: 25}
	boxCount := bundle.CountBox()
	cakeCount, appleCount := bundle.CountPartOfBox()
	response := Response(bundle, boxCount, cakeCount, appleCount)
	fmt.Println(response)
}

func Response(bundle Bundle, boxCount int, cakeCount int, appleCount int) string{
	var response string
	response += fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", bundle.Name, bundle.Cake, bundle.Apple)
	response += fmt.Sprintf("How many boxes that %v can make? %d box\n", bundle.Name, boxCount)
	response += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	return response
}