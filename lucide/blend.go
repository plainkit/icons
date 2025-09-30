package lucide

import (
	html "github.com/plainkit/html"
)

// Blend creates a Blend Lucide icon.
func Blend(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-blend", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("7")),
		html.SvgCircle(html.ACx("15"), html.ACy("15"), html.AR("7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
