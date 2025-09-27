package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowLeft creates a Square Arrow Left Lucide icon.
func SquareArrowLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m12 8-4 4 4 4"))),
		html.Child(html.SvgPath(html.AD("M16 12H8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
