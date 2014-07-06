package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// var statemap map[state][]string

var buildTests = []struct {
	in []string
}{
	{
		in: []string{"Oh, un lobo que habla, dijo caperucita"},
	},
	{
		in: []string{"En un lugar de la mancha de cuyo nombre no quiero acordarme",
			"vivia un hidalgo que se tocaba los huevos a todas horas",
		},
	},
	{
		in: []string{"En un lugar",
			"de  la mancha  vivia",
			"un  hidalgo de la leche.",
		},
	},
	{
		in: []string{},
	},
}

var addTests = []struct {
	in           string
	prefixAfter0 string
	prefixAfter1 string
}{
	// {
	// 	in:           "",
	// 	prefixAfter0: "",
	// 	prefixAfter1: "",
	// },
	{
		in:           "word",
		prefixAfter0: "",
		prefixAfter1: "word",
	},
	// {
	// 	in:           "palabra",
	// 	prefixAfter0: "word",
	// 	prefixAfter1: "palabra",
	// },
}

// func TestBuild(t *testing.T) {
// 	// for i, test := range buildTests {
// 	// nwords, nlines, _ := build(test.in)
// 	// assert.Equal(t, nlines, test.nlines, "Test nlines: %d", i)
// 	// assert.Equal(t, nwords, test.nwords, "Test nwords: %d", i)
// 	// }
// 	t.Error("TestBuild: Unimplemented")
// }

// Test the funtion add.
func TestAdd(t *testing.T) {

	//var prefix state // state has the "" string
	statemap = make(map[state][]string)

	for i, test := range addTests {
		//var prefixOrig0, prefixOrig1 string
		//prefixOrig0, prefixOrig1 = prefix.pref[0], prefix.pref[1]

		add(test.in)

		assert.Equal(t, test.prefixAfter0, prefix.pref[0], "Test for pref[0]", i)
		assert.Equal(t, test.prefixAfter1, prefix.pref[1], "Test for pref[1]", i)
		// assert.NotEqual(t, prefixOrig0, prefix.pref[0], "prefix.pref[0] has not changed",
		// 	prefix) // depends on the test..
		// assert.NotEqual(t, prefixOrig1, prefix.pref[1], "prefix.pref[1] has not changed",
		// 	prefix) // depends on the test..
		// assert.IsType(t, string, prefixOrig0, "prefixOrig is not string")
	}
	// assert.NotEmpty(t, statemap, "Statemap not empty")
}
