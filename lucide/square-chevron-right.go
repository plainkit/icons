package lucide

import (
	html "github.com/plainkit/html"
)

// SquareChevronRight creates a Square Chevron Right Lucide icon.
func SquareChevronRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-chevron-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m10 8 4 4-4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
