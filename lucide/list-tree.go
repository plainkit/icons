package lucide

import (
	html "github.com/plainkit/html"
)

// ListTree creates a List Tree Lucide icon.
func ListTree(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-tree", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 5h13"))),
		html.Child(html.SvgPath(html.AD("M13 12h8"))),
		html.Child(html.SvgPath(html.AD("M13 19h8"))),
		html.Child(html.SvgPath(html.AD("M3 10a2 2 0 0 0 2 2h3"))),
		html.Child(html.SvgPath(html.AD("M3 5v12a2 2 0 0 0 2 2h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
