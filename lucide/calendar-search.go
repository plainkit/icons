package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarSearch creates a Calendar Search Lucide icon.
func CalendarSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-search", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M21 11.75V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.25"))),
		html.Child(html.SvgPath(html.AD("m22 22-1.875-1.875"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
