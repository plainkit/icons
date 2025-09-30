package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarX creates a Calendar X Lucide icon.
func CalendarX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("m14 14-4 4")),
		html.SvgPath(html.AD("m10 14 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
