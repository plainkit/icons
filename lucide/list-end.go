package lucide

import (
	html "github.com/plainkit/html"
)

// ListEnd creates a List End Lucide icon.
func ListEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 5H3"))),
		html.Child(html.SvgPath(html.AD("M16 12H3"))),
		html.Child(html.SvgPath(html.AD("M9 19H3"))),
		html.Child(html.SvgPath(html.AD("m16 16-3 3 3 3"))),
		html.Child(html.SvgPath(html.AD("M21 5v12a2 2 0 0 1-2 2h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
