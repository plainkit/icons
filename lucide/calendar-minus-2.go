package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarMinus2 creates a Calendar Minus 2 Lucide icon.
func CalendarMinus2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-minus-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M10 16h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
