package lucide

import (
	html "github.com/plainkit/html"
)

// Lollipop creates a Lollipop Lucide icon.
func Lollipop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lollipop", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("11"), html.ACy("11"), html.AR("8"))),
		html.Child(html.SvgPath(html.AD("m21 21-4.3-4.3"))),
		html.Child(html.SvgPath(html.AD("M11 11a2 2 0 0 0 4 0 4 4 0 0 0-8 0 6 6 0 0 0 12 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
