package lucide

import (
	html "github.com/plainkit/html"
)

// CircleChevronUp creates a Circle Chevron Up Lucide icon.
func CircleChevronUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-chevron-up", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("m8 14 4-4 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
