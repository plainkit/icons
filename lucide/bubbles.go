package lucide

import (
	html "github.com/plainkit/html"
)

// Bubbles creates a Bubbles Lucide icon.
func Bubbles(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bubbles", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7.2 14.8a2 2 0 0 1 2 2"))),
		html.Child(html.SvgCircle(html.ACx("18.5"), html.ACy("8.5"), html.AR("3.5"))),
		html.Child(html.SvgCircle(html.ACx("7.5"), html.ACy("16.5"), html.AR("5.5"))),
		html.Child(html.SvgCircle(html.ACx("7.5"), html.ACy("4.5"), html.AR("2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
