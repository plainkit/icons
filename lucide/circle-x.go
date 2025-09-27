package lucide

import (
	html "github.com/plainkit/html"
)

// CircleX creates a Circle X Lucide icon.
func CircleX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m15 9-6 6"))),
		html.Child(html.SvgPath(html.AD("m9 9 6 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
