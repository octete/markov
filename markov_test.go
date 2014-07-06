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

/*
We are passing:
"This is Sparta. And this is Seville"
Which should give us something like":
{
    {
        pref: {"", "This"},
    }:  {"is"},
    {
        pref: {"This", "is"},
    }:  {"Sparta.", "Seville."},
    {
        pref: {"is", "Sparta."},
    }:  {"And"},
    {
        pref: {"Sparta.", "And"},
    }:  {"This"},
    {
        pref: {"And", "This"},
    }:  {"is"},
    {}: {"This"},
}
*/
var addTests = []struct {
	in           string
	prefixAfter0 string
	prefixAfter1 string
}{
	{
		in:           "This",
		prefixAfter0: "",
		prefixAfter1: "This",
	},
	{
		in:           "is",
		prefixAfter0: "This",
		prefixAfter1: "is",
	},
	{
		in:           "Sparta.",
		prefixAfter0: "is",
		prefixAfter1: "Sparta.",
	},
	{
		in:           "And",
		prefixAfter0: "Sparta.",
		prefixAfter1: "And",
	},
	{
		in:           "This",
		prefixAfter0: "And",
		prefixAfter1: "This",
	},
	{
		in:           "is",
		prefixAfter0: "This",
		prefixAfter1: "is",
	},
	{
		in:           "Seville.",
		prefixAfter0: "is",
		prefixAfter1: "Seville.",
	},
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
	assert.NotEmpty(t, statemap, "Statemap not empty")
	/* statemap should have now all the words for these pairs:
	And, This
	"", This
	This, is
	is, Sparta.
	Sparta., And
	*/
	// assert.NotNil(t, statemap[state{pref: "This", "is"}], "sitemap doesn't have 'This' 'is'")
	// TODO: Add all different checks here.

}
