package Hello

import "testing"

func TestFind(t *testing.T)  {
	t.Run("Hello world",func(t *testing.T) {
		want := true
		got := Find("Hello world")
		if got != want {
			t.Errorf("\ngot %v\n want %v",got,want)
		}
	})
	t.Run("Hello Solar",func(t *testing.T) {
		want := false
		got := Find("Hello Solar System")
		if got != want {
			t.Errorf("\ngot %v\n want %v",got,want)
		}
	})
	t.Run("Hello Solar",func(t *testing.T) {
		want := false
		got := Find("Hello Solar")
		if got != want {
			t.Errorf("\ngot %v\n want %v",got,want)
		}
	})	
}
