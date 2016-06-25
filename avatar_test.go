package initialser
import (
	"testing"
)

func TestKey(t *testing.T) {
	a := NewAvatar("a")
	if "200:120:Microsoft Sans Serif:#E08F70:#ffffff:a.png" != a.Key() {
		t.Error("key cal error")
	}
}
