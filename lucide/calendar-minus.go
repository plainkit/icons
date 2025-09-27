package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarMinus creates a Calendar Minus Lucide icon.
func CalendarMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 19h6"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M21 15V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8.5"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
