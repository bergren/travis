package travis

import "testing"

func TestAdd(t *testing.T) {
	a := Add(3, 4)
	if a == 7 {
		t.Log("pass")
	} else {
		t.Error("3 + 4 should 7")
	}
}
