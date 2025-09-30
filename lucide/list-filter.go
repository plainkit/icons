package lucide

import (
	html "github.com/plainkit/html"
)

// ListFilter creates a List Filter Lucide icon.
func ListFilter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-filter", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 5h20")),
		html.SvgPath(html.AD("M6 12h12")),
		html.SvgPath(html.AD("M9 19h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
