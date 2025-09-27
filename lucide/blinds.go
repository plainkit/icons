package lucide

import (
	html "github.com/plainkit/html"
)

// Blinds creates a Blinds Lucide icon.
func Blinds(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-blinds", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3h18"))),
		html.Child(html.SvgPath(html.AD("M20 7H8"))),
		html.Child(html.SvgPath(html.AD("M20 11H8"))),
		html.Child(html.SvgPath(html.AD("M10 19h10"))),
		html.Child(html.SvgPath(html.AD("M8 15h12"))),
		html.Child(html.SvgPath(html.AD("M4 3v14"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("19"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
