package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowOutUpLeft creates a Square Arrow Out Up Left Lucide icon.
func SquareArrowOutUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-out-up-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-6"))),
		html.Child(html.SvgPath(html.AD("m3 3 9 9"))),
		html.Child(html.SvgPath(html.AD("M3 9V3h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
