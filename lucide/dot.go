package lucide

import (
	html "github.com/plainkit/html"
)

// Dot creates a Dot Lucide icon.
func Dot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dot", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12.1"), html.ACy("12.1"), html.AR("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
