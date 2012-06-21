package ppm

import "testing"

func TestMakeImage(t *testing.T) {
	f := func(i, j int) Color { return Color{0, i, j} }
	img := MakeImage(23, 42, 3, 2, f)
	if img.Width != 23 {
		t.Errorf("img.Width should be 23, is %d", img.Width)
	}
	if img.Height != 42 {
		t.Errorf("img.Height should be 42, is %d", img.Height)
	}
	if len(img.Data) != 2 {
		t.Errorf("number of rows should be 2, is %d", len(img.Data))
	}
	expectedData := Data{
		Row{Color{0, 0, 0}, Color{0, 0, 1}, Color{0, 0, 2}},
		Row{Color{0, 1, 0}, Color{0, 1, 1}, Color{0, 1, 2}}}
	if !img.Data.Equals(expectedData) {
		t.Errorf("%v != %v", img, expectedData)
	}
}

func TestMakeValidImage(t *testing.T) {}

func TestMakeSolidImage(t *testing.T) {}

func TestMakeValidSolidImage(t *testing.T) {}

func TestMakeBlackImage(t *testing.T) {}

func TestMakeValidBlackImage(t *testing.T) {}
