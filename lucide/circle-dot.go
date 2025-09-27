package lucide

import (
	html "github.com/plainkit/html"
)

// CircleDot creates a Circle Dot Lucide icon.
func CircleDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-dot", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
