package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarArrowUp creates a Calendar Arrow Up Lucide icon.
func CalendarArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-arrow-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 18 4-4 4 4")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M18 22v-8")),
		html.SvgPath(html.AD("M21 11.343V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h9")),
		html.SvgPath(html.AD("M3 10h18")),
		html.SvgPath(html.AD("M8 2v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
