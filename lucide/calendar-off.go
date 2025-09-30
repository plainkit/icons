package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarOff creates a Calendar Off Lucide icon.
func CalendarOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4.2 4.2A2 2 0 0 0 3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 1.82-1.18")),
		html.SvgPath(html.AD("M21 15.5V6a2 2 0 0 0-2-2H9.5")),
		html.SvgPath(html.AD("M16 2v4")),
		html.SvgPath(html.AD("M3 10h7")),
		html.SvgPath(html.AD("M21 10h-5.5")),
		html.SvgPath(html.AD("m2 2 20 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
