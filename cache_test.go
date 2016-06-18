package initialser
import "testing"

func TestFontCache(t *testing.T) {
	fc := newFontCache(2)
	fc.get("Kokonor")
	fc.get("Krungthep")
	fc.get("Mishafi")
	fc.get("Impact")
	t.Log(fc.cache.Len())
	if fc.cache.Len() > 2 {
		t.Error("expected 2 ,actual ", fc.cache.Len())
	}

}