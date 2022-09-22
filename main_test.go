package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expect interface{}
	response string
	boxCount int
	cakeCount int
	appleCount int
	bundle Bundle
)

func TestCountBox(t *testing.T) {
	bundle = Bundle{Cake: 20, Apple: 25}
	expect = 5
	boxCount = bundle.CountBox()
	assert.Equal(t, expect, boxCount)

	bundle = Bundle{Cake: 25, Apple: 20}
	expect = 5
	boxCount = bundle.CountBox()
	assert.Equal(t, expect, boxCount)
}

func TestCountPartOfBox(t *testing.T) {
	bundle = Bundle{Cake: 20, Apple: 25}
	boxCount = bundle.CountBox()
	cakeCount, appleCount = bundle.CountPartOfBox()
	expect = 4
	assert.Equal(t, expect, cakeCount)
	expect = 5
	assert.Equal(t, expect, appleCount)

	bundle = Bundle{Cake: 25, Apple: 20}
	boxCount = bundle.CountBox()
	cakeCount, appleCount = bundle.CountPartOfBox()
	expect = 5
	assert.Equal(t, expect, cakeCount)
	expect = 4
	assert.Equal(t, expect, appleCount)
}

func TestResponse(t *testing.T) {
	var expect string

	bundle = Bundle{Cake: 20, Apple: 25}
	boxCount = bundle.CountBox()
	cakeCount, appleCount = bundle.CountPartOfBox()
	response = Response(bundle, boxCount, cakeCount, appleCount)
	expect = Response(bundle, boxCount, cakeCount, appleCount)
	expect = fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", bundle.Name, bundle.Cake, bundle.Apple)
	expect += fmt.Sprintf("How many boxes that %v can make? %d box\n", bundle.Name, boxCount)
	expect += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	assert.Equal(t, expect, response)
	
	bundle = Bundle{Cake: 25, Apple: 20}
	boxCount = bundle.CountBox()
	cakeCount, appleCount = bundle.CountPartOfBox()
	response = Response(bundle, boxCount, cakeCount, appleCount)
	expect = fmt.Sprintf("%v have %d cakes and %d apples. She want to bundle that cakes and apples into boxes and give them to her friends.\n", bundle.Name, bundle.Cake, bundle.Apple)
	expect += fmt.Sprintf("How many boxes that %v can make? %d box\n", bundle.Name, boxCount)
	expect += fmt.Sprintf("how many cakes and apples every box have? %d cake and %d apple in every box", cakeCount, appleCount)
	assert.Equal(t, expect, response)
}
