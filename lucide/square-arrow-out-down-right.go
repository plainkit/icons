package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowOutDownRight creates a Square Arrow Out Down Right Lucide icon.
func SquareArrowOutDownRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-out-down-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		html.Child(html.SvgPath(html.AD("m21 21-9-9"))),
		html.Child(html.SvgPath(html.AD("M21 15v6h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
