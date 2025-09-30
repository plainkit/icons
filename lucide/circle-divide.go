package lucide

import (
	html "github.com/plainkit/html"
)

// CircleDivide creates a Circle Divide Lucide icon.
func CircleDivide(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-divide", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("12"), html.AY2("12")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("16"), html.AY2("16")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("8"), html.AY2("8")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
