package lucide

import (
	html "github.com/plainkit/html"
)

// SquareEqual creates a Square Equal Lucide icon.
func SquareEqual(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-equal", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M7 10h10"))),
		html.Child(html.SvgPath(html.AD("M7 14h10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
