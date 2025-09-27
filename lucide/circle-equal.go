package lucide

import (
	html "github.com/plainkit/html"
)

// CircleEqual creates a Circle Equal Lucide icon.
func CircleEqual(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-equal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 10h10"))),
		html.Child(html.SvgPath(html.AD("M7 14h10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
