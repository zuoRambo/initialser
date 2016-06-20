package initialser
import (
	"testing"
)

func TestKey(t *testing.T) {
	a := NewAvatar("中国")
	if "87523f1bd4e3cef551f53eed149f8089" != a.Key() {
		t.Error("key cal error")
	}
}
