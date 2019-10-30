package travis

import "testing"

func TestMul(t *testing.T) {
	a := Mul(3, 4)
	if a == 12 {
		t.Log("pass")
	} else {
		t.Error("3 * 4 should 12")
	}
}
