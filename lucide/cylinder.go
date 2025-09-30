package lucide

import (
	html "github.com/plainkit/html"
)

// Cylinder creates a Cylinder Lucide icon.
func Cylinder(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cylinder", args)
	children := []html.SvgArg{
		html.SvgEllipse(html.ACx("12"), html.ACy("5"), html.ARx("9"), html.ARy("3")),
		html.SvgPath(html.AD("M3 5v14a9 3 0 0 0 18 0V5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
