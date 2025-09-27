package lucide

import (
	html "github.com/plainkit/html"
)

// SquareMinus creates a Square Minus Lucide icon.
func SquareMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
