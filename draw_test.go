package initialser

import (
	"testing"
)

func TestDraw(t *testing.T) {
	a := NewAvatar("H")
	d, err := NewDrawer(a)
	if err != nil {
		t.Error("not expected error ,", err)
	}
	_, err = d.Draw()
	if err != nil {
		t.Error("not expected error ,", err)
	}
	_, err = d.DrawToBytes()
	if err != nil {
		t.Error("not expected error ,", err)
	}
}

