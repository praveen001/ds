package arraylist

import "testing"

func TestNew(t *testing.T) {
	al := New()
	if al == nil {
		t.Errorf("New didn't create new array list")
	}
}
