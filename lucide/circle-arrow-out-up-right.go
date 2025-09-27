package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowOutUpRight creates a Circle Arrow Out Up Right Lucide icon.
func CircleArrowOutUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-out-up-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 12A10 10 0 1 1 12 2"))),
		html.Child(html.SvgPath(html.AD("M22 2 12 12"))),
		html.Child(html.SvgPath(html.AD("M16 2h6v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
