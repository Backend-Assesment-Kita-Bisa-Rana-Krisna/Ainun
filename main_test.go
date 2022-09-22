package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBox(t *testing.T) {
	cake = 20
	apple = 25
	expect := 5
	boxCount := countBox(cake, apple)
	assert.Equal(t, expect, boxCount)

	cake = 25
	apple = 20
	expect = 5
	boxCount = countBox(cake, apple)
	assert.Equal(t, expect, boxCount)
}

func TestCountPartOfBox(t *testing.T) {
	cake = 20
	apple = 25
	boxCount := countBox(cake, apple)
	totalItemCake, totalItemApple := countPartOfBox(cake, apple, boxCount)
	expectCakeItem := 4
	assert.Equal(t, expectCakeItem, totalItemCake)
	expectAppleItem := 5
	assert.Equal(t, expectAppleItem, totalItemApple)

	cake = 25
	apple = 20
	boxCount = countBox(cake, apple)
	totalItemCake, totalItemApple = countPartOfBox(cake, apple, boxCount)
	expectCakeItem = 5
	assert.Equal(t, expectCakeItem, totalItemCake)
	expectAppleItem = 4
	assert.Equal(t, expectAppleItem, totalItemApple)
}

func TestResponse(t *testing.T) {
	var expect string

	name = "Ainun"
	cake = 20
	apple = 25
	boxCount := countBox(cake, apple)
	cakeCount, appleCount := countPartOfBox(cake, apple, boxCount)
	response := Response(name, cake, apple, boxCount, cakeCount, appleCount)
	expect = fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", name, cake, apple)
	expect += fmt.Sprintf("How many boxes that %v can make? %d box\n", name, boxCount)
	expect += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	assert.Equal(t, expect, response)
	
	name = "Ainun"
	cake = 25
	apple = 20
	boxCount = countBox(cake, apple)
	cakeCount, appleCount = countPartOfBox(cake, apple, boxCount)
	response = Response(name, cake, apple, boxCount, cakeCount, appleCount)
	expect = fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", name, cake, apple)
	expect += fmt.Sprintf("How many boxes that %v can make? %d box\n", name, boxCount)
	expect += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	assert.Equal(t, expect, response)
}