package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDivide creates a Square Divide Lucide icon.
func SquareDivide(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-divide", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("12"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("16"), html.AY2("16"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("8"), html.AY2("8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
