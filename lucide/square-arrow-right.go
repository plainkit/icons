package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowRight creates a Square Arrow Right Lucide icon.
func SquareArrowRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-right", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M8 12h8")),
		html.SvgPath(html.AD("m12 16 4-4-4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
