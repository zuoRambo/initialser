package initialser

import (
	"testing"
)

func TestDraw(t *testing.T) {
	a := NewAvatar("H")
	a.FontSize = 160
	a.Font = "华文黑体"
	a.Size = 180
	a.Background = "ff00"

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

