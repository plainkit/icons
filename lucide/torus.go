package lucide

import (
	html "github.com/plainkit/html"
)

// Torus creates a Torus Lucide icon.
func Torus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-torus", args)
	children := []html.SvgArg{
		html.Child(html.SvgEllipse(html.ACx("12"), html.ACy("11"), html.ARx("3"), html.ARy("2"))),
		html.Child(html.SvgEllipse(html.ACx("12"), html.ACy("12.5"), html.ARx("10"), html.ARy("8.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
