package lucide

import (
	html "github.com/plainkit/html"
)

// TimerOff creates a Timer Off Lucide icon.
func TimerOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-timer-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 2h4"))),
		html.Child(html.SvgPath(html.AD("M4.6 11a8 8 0 0 0 1.7 8.7 8 8 0 0 0 8.7 1.7"))),
		html.Child(html.SvgPath(html.AD("M7.4 7.4a8 8 0 0 1 10.3 1 8 8 0 0 1 .9 10.2"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M12 12v-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
