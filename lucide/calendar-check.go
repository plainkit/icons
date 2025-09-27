package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarCheck creates a Calendar Check Lucide icon.
func CalendarCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("m9 16 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
