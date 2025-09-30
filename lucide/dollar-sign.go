package lucide

import (
	html "github.com/plainkit/html"
)

// DollarSign creates a Dollar Sign Lucide icon.
func DollarSign(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dollar-sign", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("2"), html.AY2("22")),
		html.SvgPath(html.AD("M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
