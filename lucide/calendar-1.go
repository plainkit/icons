package lucide

import (
	html "github.com/plainkit/html"
)

// Calendar1 creates a Calendar 1 Lucide icon.
func Calendar1(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-1", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 14h1v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
