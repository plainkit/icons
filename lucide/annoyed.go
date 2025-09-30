package lucide

import (
	html "github.com/plainkit/html"
)

// Annoyed creates a Annoyed Lucide icon.
func Annoyed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-annoyed", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M8 15h8")),
		html.SvgPath(html.AD("M8 9h2")),
		html.SvgPath(html.AD("M14 9h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
