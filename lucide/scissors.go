package lucide

import (
	html "github.com/plainkit/html"
)

// Scissors creates a Scissors Lucide icon.
func Scissors(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scissors", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M8.12 8.12 12 12"))),
		html.Child(html.SvgPath(html.AD("M20 4 8.12 15.88"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M14.8 14.8 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
