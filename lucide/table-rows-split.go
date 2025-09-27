package lucide

import (
	html "github.com/plainkit/html"
)

// TableRowsSplit creates a Table Rows Split Lucide icon.
func TableRowsSplit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-rows-split", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 10h2"))),
		html.Child(html.SvgPath(html.AD("M15 22v-8"))),
		html.Child(html.SvgPath(html.AD("M15 2v4"))),
		html.Child(html.SvgPath(html.AD("M2 10h2"))),
		html.Child(html.SvgPath(html.AD("M20 10h2"))),
		html.Child(html.SvgPath(html.AD("M3 19h18"))),
		html.Child(html.SvgPath(html.AD("M3 22v-6a2 2 135 0 1 2-2h14a2 2 45 0 1 2 2v6"))),
		html.Child(html.SvgPath(html.AD("M3 2v2a2 2 45 0 0 2 2h14a2 2 135 0 0 2-2V2"))),
		html.Child(html.SvgPath(html.AD("M8 10h2"))),
		html.Child(html.SvgPath(html.AD("M9 22v-8"))),
		html.Child(html.SvgPath(html.AD("M9 2v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
