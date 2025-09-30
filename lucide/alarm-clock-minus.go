package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmClockMinus creates a Alarm Clock Minus Lucide icon.
func AlarmClockMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-clock-minus", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("8")),
		html.SvgPath(html.AD("M5 3 2 6")),
		html.SvgPath(html.AD("m22 6-3-3")),
		html.SvgPath(html.AD("M6.38 18.7 4 21")),
		html.SvgPath(html.AD("M17.64 18.67 20 21")),
		html.SvgPath(html.AD("M9 13h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
