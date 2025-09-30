package lucide

import (
	html "github.com/plainkit/html"
)

// CircleChevronDown creates a Circle Chevron Down Lucide icon.
func CircleChevronDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-chevron-down", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("m16 10-4 4-4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
