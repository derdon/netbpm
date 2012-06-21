package ppm

import "testing"

func TestInvalidHeight(t *testing.T) {
	img := MakeBlackImage(3, 5, 3, 2)
	valid := img.IsValid()
	if valid {
		t.Errorf("%v should be invalid, but it isn't!", img)
	}
}

/*
func TestInvalidWidth(t *testing.T) {

	valid := i.isValid()
}
*/
