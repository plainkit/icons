package lucide

import (
	html "github.com/plainkit/html"
)

// Circle creates a Circle Lucide icon.
func Circle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
