package lucide

import (
	html "github.com/plainkit/html"
)

// CircleChevronLeft creates a Circle Chevron Left Lucide icon.
func CircleChevronLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-chevron-left", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("m14 16-4-4 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
