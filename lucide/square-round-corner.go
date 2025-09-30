package lucide

import (
	html "github.com/plainkit/html"
)

// SquareRoundCorner creates a Square Round Corner Lucide icon.
func SquareRoundCorner(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-round-corner", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 11a8 8 0 0 0-8-8")),
		html.SvgPath(html.AD("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
