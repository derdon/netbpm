package ppm

import "testing"

func TestDifferentLength(t *testing.T) {
	d1 := Data{
		Row{black, black, black},
		Row{black, black, black}}
	d2 := Data{
		Row{black, black, black}}
	errorOnEqualData(d1, d2, t)
}

func TestDifferentWidth(t *testing.T) {
	d1 := Data{
		Row{black, black},
		Row{black, black}}
	d2 := Data{
		Row{black, black, black},
		Row{black, black, black}}
	errorOnEqualData(d1, d2, t)
}

func TestMixedWidths(t *testing.T) {
	d1 := Data{
		Row{black, black, black},
		Row{black, black}}
	d2 := Data{
		Row{black, black, black},
		Row{black, black}}
	if !d1.Equals(d2) {
		t.Errorf("%v != %v", d1, d2)
	}
}
