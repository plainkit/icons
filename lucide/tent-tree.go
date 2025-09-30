package lucide

import (
	html "github.com/plainkit/html"
)

// TentTree creates a Tent Tree Lucide icon.
func TentTree(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tent-tree", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("4"), html.ACy("4"), html.AR("2")),
		html.SvgPath(html.AD("m14 5 3-3 3 3")),
		html.SvgPath(html.AD("m14 10 3-3 3 3")),
		html.SvgPath(html.AD("M17 14V2")),
		html.SvgPath(html.AD("M17 14H7l-5 8h20Z")),
		html.SvgPath(html.AD("M8 14v8")),
		html.SvgPath(html.AD("m9 14 5 8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
