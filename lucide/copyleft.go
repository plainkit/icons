package lucide

import (
	html "github.com/plainkit/html"
)

// Copyleft creates a Copyleft Lucide icon.
func Copyleft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-copyleft", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M9.17 14.83a4 4 0 1 0 0-5.66"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
