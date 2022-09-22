package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	apple int
	cake int
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
