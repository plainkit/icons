package lucide

import (
	html "github.com/plainkit/html"
)

// Music2 creates a Music 2 Lucide icon.
func Music2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-music-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("18"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M12 18V2l7 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
