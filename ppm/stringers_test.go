package ppm

import (
	"fmt"
	"testing"
)

func TestColorString(t *testing.T) {
	c := Color{23, 112, 240}
	s := fmt.Sprintf("%v", c)
	expectedS := "Color(23, 112, 240)"
	errorOnUnequalStrings(s, expectedS, t)
}

func TestRowString(t *testing.T) {
	r := Row{Color{1, 2, 3}, Color{23, 42, 1337}}
	s := fmt.Sprintf("%v", r)
	expectedS := "[Color(1, 2, 3), Color(23, 42, 1337)]"
	errorOnUnequalStrings(s, expectedS, t)
}

func TestDataString(t *testing.T) {
	d := Data{
		Row{Color{50, 60, 70}, Color{20, 30, 40}},
		Row{Color{13, 26, 39}}}
	s := fmt.Sprintf("%v", d)
	expectedS := "[\n" +
		"    [Color(50, 60, 70), Color(20, 30, 40)],\n" +
		"    [Color(13, 26, 39)]\n" +
		"]"
	errorOnUnequalStrings(s, expectedS, t)
}
