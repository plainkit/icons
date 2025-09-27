package lucide

import (
	html "github.com/plainkit/html"
)

// Calendar creates a Calendar Lucide icon.
func Calendar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
