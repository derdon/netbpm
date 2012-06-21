package ppm

import (
	"fmt"
	"testing"
)

func TestGetWidthAndHeightInvalidWidth(t *testing.T) {
	_, _, err := getWidthAndHeight("foo 56")
	errorOnErrIsNil(err, t)
	s := fmt.Sprintf("%s", err)
	expectedErrMsg := "invalid width: foo"
	if s != expectedErrMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedErrMsg, s)
	}
}

func TestGetWidthAndHeightInvalidHeight(t *testing.T) {
	_, _, err := getWidthAndHeight("42 bar")
	errorOnErrIsNil(err, t)
	s := fmt.Sprintf("%s", err)
	expectedErrMsg := "invalid height: bar"
	if s != expectedErrMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedErrMsg, s)
	}
}

func TestGetWidthAndHeightValid(t *testing.T) {
	width, height, err := getWidthAndHeight("42 10")
	errorOnErrIsNotNil(err, t)
	if width != 42 {
		t.Errorf("width != 42, width == %v", width)
	}
	if height != 10 {
		t.Errorf("height != 10, height == %v", height)
	}
}

func TestGetWidthAndHeightMissingHeight(t *testing.T) {
	// it doesn't matter if there's a number or simply any string on this line
	// if there is one declaration only.
	_, _, err := getWidthAndHeight("foobar")
	errorOnErrIsNil(err, t)
	s := fmt.Sprintf("%s", err)
	expectedErrMsg := "missing information: height"
	if s != expectedErrMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedErrMsg, s)
	}
}
