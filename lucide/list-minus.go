package lucide

import (
	html "github.com/plainkit/html"
)

// ListMinus creates a List Minus Lucide icon.
func ListMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 5H3"))),
		html.Child(html.SvgPath(html.AD("M11 12H3"))),
		html.Child(html.SvgPath(html.AD("M16 19H3"))),
		html.Child(html.SvgPath(html.AD("M21 12h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
