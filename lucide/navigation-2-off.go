package lucide

import (
	html "github.com/plainkit/html"
)

// Navigation2Off creates a Navigation 2 Off Lucide icon.
func Navigation2Off(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-navigation-2-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9.31 9.31 5 21l7-4 7 4-1.17-3.17"))),
		html.Child(html.SvgPath(html.AD("M14.53 8.88 12 2l-1.17 3.17"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
