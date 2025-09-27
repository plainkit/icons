package lucide

import (
	html "github.com/plainkit/html"
)

// ListIndentIncrease creates a List Indent Increase Lucide icon.
func ListIndentIncrease(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-indent-increase", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 5H11"))),
		html.Child(html.SvgPath(html.AD("M21 12H11"))),
		html.Child(html.SvgPath(html.AD("M21 19H11"))),
		html.Child(html.SvgPath(html.AD("m3 8 4 4-4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
