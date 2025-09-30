package lucide

import (
	html "github.com/plainkit/html"
)

// CornerUpLeft creates a Corner Up Left Lucide icon.
func CornerUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-corner-up-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 20v-7a4 4 0 0 0-4-4H4")),
		html.SvgPath(html.AD("M9 14 4 9l5-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
