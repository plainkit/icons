package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowOutDownLeft creates a Circle Arrow Out Down Left Lucide icon.
func CircleArrowOutDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-out-down-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 12a10 10 0 1 1 10 10")),
		html.SvgPath(html.AD("m2 22 10-10")),
		html.SvgPath(html.AD("M8 22H2v-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
