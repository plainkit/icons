package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarSync creates a Calendar Sync Lucide icon.
func CalendarSync(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-sync", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 10v4h4"))),
		html.Child(html.SvgPath(html.AD("m11 14 1.535-1.605a5 5 0 0 1 8 1.5"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("m21 18-1.535 1.605a5 5 0 0 1-8-1.5"))),
		html.Child(html.SvgPath(html.AD("M21 22v-4h-4"))),
		html.Child(html.SvgPath(html.AD("M21 8.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h4.3"))),
		html.Child(html.SvgPath(html.AD("M3 10h4"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
