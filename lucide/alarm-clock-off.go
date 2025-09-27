package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmClockOff creates a Alarm Clock Off Lucide icon.
func AlarmClockOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-clock-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6.87 6.87a8 8 0 1 0 11.26 11.26"))),
		html.Child(html.SvgPath(html.AD("M19.9 14.25a8 8 0 0 0-9.15-9.15"))),
		html.Child(html.SvgPath(html.AD("m22 6-3-3"))),
		html.Child(html.SvgPath(html.AD("M6.26 18.67 4 21"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M4 4 2 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
