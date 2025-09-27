package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowDownRight creates a Square Arrow Down Right Lucide icon.
func SquareArrowDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-down-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m8 8 8 8"))),
		html.Child(html.SvgPath(html.AD("M16 8v8H8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
