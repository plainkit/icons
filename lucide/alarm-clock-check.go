package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmClockCheck creates a Alarm Clock Check Lucide icon.
func AlarmClockCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-clock-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("8"))),
		html.Child(html.SvgPath(html.AD("M5 3 2 6"))),
		html.Child(html.SvgPath(html.AD("m22 6-3-3"))),
		html.Child(html.SvgPath(html.AD("M6.38 18.7 4 21"))),
		html.Child(html.SvgPath(html.AD("M17.64 18.67 20 21"))),
		html.Child(html.SvgPath(html.AD("m9 13 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
