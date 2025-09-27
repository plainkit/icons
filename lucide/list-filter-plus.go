package lucide

import (
	html "github.com/plainkit/html"
)

// ListFilterPlus creates a List Filter Plus Lucide icon.
func ListFilterPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-filter-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 5H2"))),
		html.Child(html.SvgPath(html.AD("M6 12h12"))),
		html.Child(html.SvgPath(html.AD("M9 19h6"))),
		html.Child(html.SvgPath(html.AD("M16 5h6"))),
		html.Child(html.SvgPath(html.AD("M19 8V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
