package lucide

import (
	html "github.com/plainkit/html"
)

// Divide creates a Divide Lucide icon.
func Divide(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-divide", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("6"), html.AR("1")),
		html.SvgLine(html.AX1("5"), html.AX2("19"), html.AY1("12"), html.AY2("12")),
		html.SvgCircle(html.ACx("12"), html.ACy("18"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
