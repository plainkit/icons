package lucide

import (
	html "github.com/plainkit/html"
)

// SquareCheckBig creates a Square Check Big Lucide icon.
func SquareCheckBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-check-big", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 10.656V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h12.344"))),
		html.Child(html.SvgPath(html.AD("m9 11 3 3L22 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
