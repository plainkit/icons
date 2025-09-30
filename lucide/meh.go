package lucide

import (
	html "github.com/plainkit/html"
)

// Meh creates a Meh Lucide icon.
func Meh(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-meh", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("15"), html.AY2("15")),
		html.SvgLine(html.AX1("9"), html.AX2("9.01"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("9"), html.AY2("9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
