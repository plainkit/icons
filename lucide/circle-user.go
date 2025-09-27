package lucide

import (
	html "github.com/plainkit/html"
)

// CircleUser creates a Circle User Lucide icon.
func CircleUser(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-user", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M7 20.662V19a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.662"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
