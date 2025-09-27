package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarPlus creates a Calendar Plus Lucide icon.
func CalendarPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 19h6"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M19 16v6"))),
		html.Child(html.SvgPath(html.AD("M21 12.598V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8.5"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
