package lucide

import (
	html "github.com/plainkit/html"
)

// NavigationOff creates a Navigation Off Lucide icon.
func NavigationOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-navigation-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8.43 8.43 3 11l8 2 2 8 2.57-5.43")),
		html.SvgPath(html.AD("M17.39 11.73 22 2l-9.73 4.61")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
