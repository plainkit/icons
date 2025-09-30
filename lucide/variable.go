package lucide

import (
	html "github.com/plainkit/html"
)

// Variable creates a Variable Lucide icon.
func Variable(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-variable", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 21s-4-3-4-9 4-9 4-9")),
		html.SvgPath(html.AD("M16 3s4 3 4 9-4 9-4 9")),
		html.SvgLine(html.AX1("15"), html.AX2("9"), html.AY1("9"), html.AY2("15")),
		html.SvgLine(html.AX1("9"), html.AX2("15"), html.AY1("9"), html.AY2("15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
