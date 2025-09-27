package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowOutUpRight creates a Square Arrow Out Up Right Lucide icon.
func SquareArrowOutUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-out-up-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6"))),
		html.Child(html.SvgPath(html.AD("m21 3-9 9"))),
		html.Child(html.SvgPath(html.AD("M15 3h6v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
