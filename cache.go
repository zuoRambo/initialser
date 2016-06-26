package initialser
import (
	"sync"
	"github.com/golang/freetype/truetype"
	"github.com/leonlau/fonts"
)
var (
	_font_cache = newFontCache(20)
)
type fontCache struct {
	rw    sync.RWMutex
	cache *LRUCache
}
//AppendFontPath append font search path
func AppendFontPath(path string) {
	fonts.AppendPath(path)
}

func OnlyPath(path string) {
	fonts.OnlyPath(path)
}

func newFontCache(max int) fontCache {
	return fontCache{
		cache:NewLRUCache(max),
	}
}
func (f fontCache)get(fontName string) (*truetype.Font, error) {
	if val, ok := f.cache.Get(fontName); ok {
		return val.(*truetype.Font), nil
	}else {
		f.rw.Lock();
		defer f.rw.Unlock();
		if val, ok = f.cache.Get(fontName); ok {
			return val.(*truetype.Font), nil
		}
		fontBytes, err := fonts.LoadFont(fontName)
		if err != nil {
			return nil, err
		}
		font, err := truetype.Parse(fontBytes)
		if err != nil {
			return nil, err
		}
		f.cache.Set(fontName, font)
		return font, nil
	}

}