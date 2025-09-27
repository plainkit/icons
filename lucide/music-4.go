package lucide

import (
	html "github.com/plainkit/html"
)

// Music4 creates a Music 4 Lucide icon.
func Music4(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-music-4", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 18V5l12-2v13"))),
		html.Child(html.SvgPath(html.AD("m9 9 12-2"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("18"), html.AR("3"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("16"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
