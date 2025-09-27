package lucide

import (
	html "github.com/plainkit/html"
)

// TimerReset creates a Timer Reset Lucide icon.
func TimerReset(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-timer-reset", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 2h4"))),
		html.Child(html.SvgPath(html.AD("M12 14v-4"))),
		html.Child(html.SvgPath(html.AD("M4 13a8 8 0 0 1 8-7 8 8 0 1 1-5.3 14L4 17.6"))),
		html.Child(html.SvgPath(html.AD("M9 17H4v5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
