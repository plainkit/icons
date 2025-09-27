package lucide

import (
	html "github.com/plainkit/html"
)

// Target creates a Target Lucide icon.
func Target(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-target", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("6"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
