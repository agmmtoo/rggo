package add

import "testing"

func TestAdd(t *testing.T) {
	a := 2
	b := 3
	exp := 5
	res := add(a, b)

	if exp != res {
		t.Errorf("Exptected %d, go %d.", exp, res)
	}
}
