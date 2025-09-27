package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarRange creates a Calendar Range Lucide icon.
func CalendarRange(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-range", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M17 14h-6"))),
		html.Child(html.SvgPath(html.AD("M13 18H7"))),
		html.Child(html.SvgPath(html.AD("M7 14h.01"))),
		html.Child(html.SvgPath(html.AD("M17 18h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
