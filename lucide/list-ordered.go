package lucide

import (
	html "github.com/plainkit/html"
)

// ListOrdered creates a List Ordered Lucide icon.
func ListOrdered(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-ordered", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 5h10"))),
		html.Child(html.SvgPath(html.AD("M11 12h10"))),
		html.Child(html.SvgPath(html.AD("M11 19h10"))),
		html.Child(html.SvgPath(html.AD("M4 4h1v5"))),
		html.Child(html.SvgPath(html.AD("M4 9h2"))),
		html.Child(html.SvgPath(html.AD("M6.5 20H3.4c0-1 2.6-1.925 2.6-3.5a1.5 1.5 0 0 0-2.6-1.02"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
