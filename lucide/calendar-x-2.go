package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarX2 creates a Calendar X 2 Lucide icon.
func CalendarX2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-x-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("m17 22 5-5")),
		html.SvgPath(html.AD("m17 17 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
