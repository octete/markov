package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var buildTests = []struct {
	in     []string
	nwords int
	nlines int
}{
	{
		in:     []string{"Oh, un lobo que habla, dijo caperucita"},
		nwords: 7,
		nlines: 1,
	},
	{
		in: []string{"En un lugar de la mancha de cuyo nombre no quiero acordarme",
			"vivia un hidalgo que se tocaba los huevos a todas horas",
		},
		nwords: 23,
		nlines: 2,
	},
	{
		in: []string{"En un lugar",
			"de  la mancha  vivia",
			"un  hidalgo de la leche.",
		},
		nwords: 12,
		nlines: 3,
	},
	{
		in:     []string{},
		nwords: 0,
		nlines: 0,
	},
}

func TestBuild(t *testing.T) {
	for i, test := range buildTests {
		nwords, nlines, _ := build(test.in)
		assert.Equal(t, nlines, test.nlines, "Test nlines: %d", i)
		assert.Equal(t, nwords, test.nwords, "Test nwords: %d", i)
	}
}
