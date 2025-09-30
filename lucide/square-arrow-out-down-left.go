package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowOutDownLeft creates a Square Arrow Out Down Left Lucide icon.
func SquareArrowOutDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-out-down-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 21h6a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6")),
		html.SvgPath(html.AD("m3 21 9-9")),
		html.SvgPath(html.AD("M9 21H3v-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
