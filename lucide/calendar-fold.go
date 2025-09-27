package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarFold creates a Calendar Fold Lucide icon.
func CalendarFold(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-fold", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M21 17V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h11Z"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M15 22v-4a2 2 0 0 1 2-2h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
