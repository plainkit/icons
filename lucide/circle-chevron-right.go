package lucide

import (
	html "github.com/plainkit/html"
)

// CircleChevronRight creates a Circle Chevron Right Lucide icon.
func CircleChevronRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-chevron-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m10 8 4 4-4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
