package lucide

import (
	html "github.com/plainkit/html"
)

// ListCollapse creates a List Collapse Lucide icon.
func ListCollapse(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-collapse", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 5h11"))),
		html.Child(html.SvgPath(html.AD("M10 12h11"))),
		html.Child(html.SvgPath(html.AD("M10 19h11"))),
		html.Child(html.SvgPath(html.AD("m3 10 3-3-3-3"))),
		html.Child(html.SvgPath(html.AD("m3 20 3-3-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
