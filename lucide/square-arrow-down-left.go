package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowDownLeft creates a Square Arrow Down Left Lucide icon.
func SquareArrowDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-down-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m16 8-8 8"))),
		html.Child(html.SvgPath(html.AD("M16 16H8V8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
