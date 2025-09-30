package lucide

import (
	html "github.com/plainkit/html"
)

// SquareChevronLeft creates a Square Chevron Left Lucide icon.
func SquareChevronLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-chevron-left", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("m14 16-4-4 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
