package add

import "testing"

func TestAdd(t *testing.T) {
	a := 2
	b := 3
	exp := 5

	result := add(a, b)

	if exp != result {
		t.Errorf("Expected %d, got %d.", exp, result)
	}
}
