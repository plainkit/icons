package lucide

import (
	html "github.com/plainkit/html"
)

// Cone creates a Cone Lucide icon.
func Cone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cone", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m20.9 18.55-8-15.98a1 1 0 0 0-1.8 0l-8 15.98")),
		html.SvgEllipse(html.ACx("12"), html.ACy("19"), html.ARx("9"), html.ARy("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
