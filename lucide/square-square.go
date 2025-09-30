package lucide

import (
	html "github.com/plainkit/html"
)

// SquareSquare creates a Square Square Lucide icon.
func SquareSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-square", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("8"), html.AY("8"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
