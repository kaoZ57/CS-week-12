package Hello

import "testing"

func TestSum(t *testing.T)   {
	want := 1
	got := Sum(1, 1)
	if got != want {
		t.Error("\nsum 1,1 \n", want)
	}
	want = 1
	got = Sum(2, 1)
	if got != want {
		t.Error("\nsum 2,1 \n", want)
	}
	want = 1
	got = Sum(2, 4)
	if got != want {
		t.Error("\nsum 2,4 \n", want)
	}
}
