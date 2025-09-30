package lucide

import (
	html "github.com/plainkit/html"
)

// SquareSplitHorizontal creates a Square Split Horizontal Lucide icon.
func SquareSplitHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-split-horizontal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 19H5c-1 0-2-1-2-2V7c0-1 1-2 2-2h3")),
		html.SvgPath(html.AD("M16 5h3c1 0 2 1 2 2v10c0 1-1 2-2 2h-3")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("4"), html.AY2("20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
