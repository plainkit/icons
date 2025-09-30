package lucide

import (
	html "github.com/plainkit/html"
)

// AlarmClockPlus creates a Alarm Clock Plus Lucide icon.
func AlarmClockPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-alarm-clock-plus", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("8")),
		html.SvgPath(html.AD("M5 3 2 6")),
		html.SvgPath(html.AD("m22 6-3-3")),
		html.SvgPath(html.AD("M6.38 18.7 4 21")),
		html.SvgPath(html.AD("M17.64 18.67 20 21")),
		html.SvgPath(html.AD("M12 10v6")),
		html.SvgPath(html.AD("M9 13h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
