package lucide

import (
	html "github.com/plainkit/html"
)

// SquareSlash creates a Square Slash Lucide icon.
func SquareSlash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-slash", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgLine(html.AX1("9"), html.AX2("15"), html.AY1("15"), html.AY2("9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
