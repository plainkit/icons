package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarArrowDown creates a Calendar Arrow Down Lucide icon.
func CalendarArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-arrow-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 18 4 4 4-4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M18 14v8")),
		html.SvgPath(html.AD("M21 11.354V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.343")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("M8 2v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
