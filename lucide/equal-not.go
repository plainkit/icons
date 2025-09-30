package lucide

import (
	html "github.com/plainkit/html"
)

// EqualNot creates a Equal Not Lucide icon.
func EqualNot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-equal-not", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("5"), html.AX2("19"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("5"), html.AX2("19"), html.AY1("15"), html.AY2("15")),
		html.SvgLine(html.AX1("19"), html.AX2("5"), html.AY1("5"), html.AY2("19")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
