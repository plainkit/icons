package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarClock creates a Calendar Clock Lucide icon.
func CalendarClock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-clock", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 14v2.2l1.6 1")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M21 7.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.5")),
		html.SvgPath(html.AD("M3 10h5")),
		html.SvgPath(html.AD("M8 2v4")),
		html.SvgCircle(html.ACx("16"), html.ACy("16"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
