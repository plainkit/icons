package lucide

import (
	html "github.com/plainkit/html"
)

// CircleSmall creates a Circle Small Lucide icon.
func CircleSmall(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-small", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
