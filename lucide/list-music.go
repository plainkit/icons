package lucide

import (
	html "github.com/plainkit/html"
)

// ListMusic creates a List Music Lucide icon.
func ListMusic(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-music", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 5H3")),
		html.SvgPath(html.AD("M11 12H3")),
		html.SvgPath(html.AD("M11 19H3")),
		html.SvgPath(html.AD("M21 16V5")),
		html.SvgCircle(html.ACx("18"), html.ACy("16"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
