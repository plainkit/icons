package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmClock creates a Alarm Clock Lucide icon.
func AlarmClock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-clock", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("8")),
		html.SvgPath(html.AD("M12 9v4l2 2")),
		html.SvgPath(html.AD("M5 3 2 6")),
		html.SvgPath(html.AD("m22 6-3-3")),
		html.SvgPath(html.AD("M6.38 18.7 4 21")),
		html.SvgPath(html.AD("M17.64 18.67 20 21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
