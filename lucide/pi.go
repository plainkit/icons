package lucide

import (
	html "github.com/plainkit/html"
)

// Pi creates a Pi Lucide icon.
func Pi(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pi", args)
	children := []html.SvgArg{
		html.Child(html.SvgLine(html.AX1("9"), html.AX2("9"), html.AY1("4"), html.AY2("20"))),
		html.Child(html.SvgPath(html.AD("M4 7c0-1.7 1.3-3 3-3h13"))),
		html.Child(html.SvgPath(html.AD("M18 20c-1.7 0-3-1.3-3-3V4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
