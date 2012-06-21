package ppm

import "testing"

func errorOnErrIsNil(err error, t *testing.T) {
	if err == nil {
		t.Error("expected and error, but err is nil")
	}
}

func errorOnErrIsNotNil(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expected err to be nil, err is %v", err)
	}
}

func errorOnEqualData(d1, d2 Data, t *testing.T) {
	if d1.Equals(d2) {
		t.Errorf("%v == %v, although they shouldn't be equal", d1, d2)
	}
}

func errorOnUnequalStrings(s1, s2 string, t *testing.T) {
	if s1 != s2 {
		t.Errorf("expected '%s' as output, got '%s'", s1, s2)
	}
}
