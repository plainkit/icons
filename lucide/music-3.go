package lucide

import (
	html "github.com/plainkit/html"
)

// Music3 creates a Music 3 Lucide icon.
func Music3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-music-3", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("18"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M16 18V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
