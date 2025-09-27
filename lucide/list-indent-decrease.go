package lucide

import (
	html "github.com/plainkit/html"
)

// ListIndentDecrease creates a List Indent Decrease Lucide icon.
func ListIndentDecrease(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-indent-decrease", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 5H11"))),
		html.Child(html.SvgPath(html.AD("M21 12H11"))),
		html.Child(html.SvgPath(html.AD("M21 19H11"))),
		html.Child(html.SvgPath(html.AD("m7 8-4 4 4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
