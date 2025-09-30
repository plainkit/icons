package lucide

import (
	html "github.com/plainkit/html"
)

// SquareChevronUp creates a Square Chevron Up Lucide icon.
func SquareChevronUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-chevron-up", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("m8 14 4-4 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
