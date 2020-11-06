package main

import (
	"github.com/gdamore/tcell"
)

var (
	// stylesPalette = stylesWikipedia
	stylesPalette = stylesGolly
	// stylesPalette = stylesChander1
	// stylesPalette = stylesChander2
)

var (
	stylesWikipedia = []tcell.Style{
		empty:        styleByHexColor(0x000000), // black
		electronHead: styleByHexColor(0x0281ff), // blue
		electronTail: styleByHexColor(0xff3e01), // red
		conductor:    styleByHexColor(0xffd803), // yellow
	}

	stylesGolly = []tcell.Style{
		empty:        styleByHexColor(0x3f3f3f), // black
		electronHead: styleByHexColor(0x4098f7), // blue
		electronTail: styleByHexColor(0xffffff), // white
		conductor:    styleByHexColor(0xee9639), // orange
	}

	stylesChander1 = []tcell.Style{
		empty:        styleByHexColor(0x005637),
		electronHead: styleByHexColor(0xffffff),
		electronTail: styleByHexColor(0xe6d5a9),
		conductor:    styleByHexColor(0xceb573),
	}

	stylesChander2 = []tcell.Style{
		empty:        styleByHexColor(0x003718),
		electronHead: styleByHexColor(0xffffff),
		electronTail: styleByHexColor(0xd4d696),
		conductor:    styleByHexColor(0xacaf58),
	}
)

func styleByHexColor(color int32) tcell.Style {
	c := tcell.NewHexColor(color)
	return tcell.StyleDefault.
		Background(c).
		Foreground(c)
}
