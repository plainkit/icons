package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarCheck2 creates a Calendar Check 2 Lucide icon.
func CalendarCheck2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-check-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M21 14V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("m16 20 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
