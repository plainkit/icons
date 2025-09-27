package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarDays creates a Calendar Days Lucide icon.
func CalendarDays(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-days", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 14h.01"))),
		html.Child(html.SvgPath(html.AD("M12 14h.01"))),
		html.Child(html.SvgPath(html.AD("M16 14h.01"))),
		html.Child(html.SvgPath(html.AD("M8 18h.01"))),
		html.Child(html.SvgPath(html.AD("M12 18h.01"))),
		html.Child(html.SvgPath(html.AD("M16 18h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
