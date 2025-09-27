package lucide

import (
	html "github.com/plainkit/html"
)

// SquareCode creates a Square Code Lucide icon.
func SquareCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-code", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10 9-3 3 3 3"))),
		html.Child(html.SvgPath(html.AD("m14 15 3-3-3-3"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
