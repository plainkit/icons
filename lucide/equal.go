package lucide

import (
	html "github.com/plainkit/html"
)

// Equal creates a Equal Lucide icon.
func Equal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-equal", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("5"), html.AX2("19"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("5"), html.AX2("19"), html.AY1("15"), html.AY2("15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
