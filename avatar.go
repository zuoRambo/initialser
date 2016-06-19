package initialser

import (
	"fmt"
	"strings"
)

//Avatar avatar
type Avatar struct {
	Size       int
	FontSize   int
	Font       string
	Background string
	Color      string
	text       string
	initial    rune
}

const svgTpl = `
<svg xmlns="http://www.w3.org/2000/svg" width="%dpx" height="%dpx">
      <g>
        <rect x="0" y="0" fill="%s" width="%dpx" height="%dpx">
        </rect>
        <text y="50%%" x="50%%" fill="%s"
              text-anchor="middle" dominant-baseline="central"
              style="font-family: %s; font-size: %dpx">
          %s
        </text>
      </g>
    </svg>
`

//NewAvatar Avatar
func NewAvatar(text string) *Avatar {
	return &Avatar{
		text:text,
		Size: 200,
		initial:[]rune(strings.ToUpper(text))[0],
		FontSize:120,
		Color:      "#ffffff",
		Background: "#E08F70",
		Font:       "sans-serif",
	}
}

func (a *Avatar)Text(text string) *Avatar {
	a.text = text;
	a.initial = []rune(strings.ToUpper(text))[0]
	return a;
}


//Valid check params
func (a *Avatar) Valid() bool {
	//TODO check params valid
	return true
}
//Key cache key
func (a *Avatar) Key() string {
	return fmt.Sprintf("%d_%d_%s_%s_%s_%d",
		a.Size,
		a.FontSize,
		a.Font,
		a.Background,
		a.Color,
		a.initial)
}

//Svg format svg
func (a *Avatar) Svg() string {
	c := a.Color
	if len(c) == 6 {
		c = "#" + c
	}
	bg := a.Background
	if len(bg) == 6 {
		bg = "#" + bg
	}
	return fmt.Sprintf(svgTpl, a.Size, a.Size, bg, a.Size, a.Size, c, a.Font, a.FontSize, string(a.initial))
}
