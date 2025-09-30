package lucide

import (
	html "github.com/plainkit/html"
)

// SquareChevronDown creates a Square Chevron Down Lucide icon.
func SquareChevronDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-chevron-down", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("m16 10-4 4-4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
