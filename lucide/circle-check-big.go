package lucide

import (
	html "github.com/plainkit/html"
)

// CircleCheckBig creates a Circle Check Big Lucide icon.
func CircleCheckBig(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-check-big", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21.801 10A10 10 0 1 1 17 3.335"))),
		html.Child(html.SvgPath(html.AD("m9 11 3 3L22 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
