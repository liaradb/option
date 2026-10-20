package option

import "testing"

func TestOption_Some(t *testing.T) {
	t.Parallel()

	s := Some(1)
	if v, ok := s.Value(); !ok {
		t.Error("should have value")
	} else if v != 1 {
		t.Errorf("incorrect value: %v, expected: %v", v, 1)
	}
}

func TestOption_None(t *testing.T) {
	t.Parallel()

	n := None[int]()
	if v, ok := n.Value(); ok {
		t.Error("should not have value")
	} else if v != 0 {
		t.Errorf("incorrect value: %v, expected: %v", v, 0)
	}
}
