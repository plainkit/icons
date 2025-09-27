package lucide

import (
	html "github.com/plainkit/html"
)

// CircleSlash creates a Circle Slash Lucide icon.
func CircleSlash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-slash", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgLine(html.AX1("9"), html.AX2("15"), html.AY1("15"), html.AY2("9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
