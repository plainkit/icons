package lucide

import (
	html "github.com/plainkit/html"
)

// SquareSplitVertical creates a Square Split Vertical Lucide icon.
func SquareSplitVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-split-vertical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 8V5c0-1 1-2 2-2h10c1 0 2 1 2 2v3")),
		html.SvgPath(html.AD("M19 16v3c0 1-1 2-2 2H7c-1 0-2-1-2-2v-3")),
		html.SvgLine(html.AX1("4"), html.AX2("20"), html.AY1("12"), html.AY2("12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
