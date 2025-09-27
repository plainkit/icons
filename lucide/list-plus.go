package lucide

import (
	html "github.com/plainkit/html"
)

// ListPlus creates a List Plus Lucide icon.
func ListPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 5H3"))),
		html.Child(html.SvgPath(html.AD("M11 12H3"))),
		html.Child(html.SvgPath(html.AD("M16 19H3"))),
		html.Child(html.SvgPath(html.AD("M18 9v6"))),
		html.Child(html.SvgPath(html.AD("M21 12h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
