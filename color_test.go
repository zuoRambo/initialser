package initialser

import (
	"testing"
)

func TestHexToRGBA(t *testing.T) {
	_, err := hexToRGBA("#345b38")
	if err != nil {
		t.Error("not expected error,", err)
	}
	_, err = hexToRGBA("#f0c")
	if err != nil {
		t.Error("not expected error,", err)
	}
}




