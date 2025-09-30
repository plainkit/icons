package lucide

import (
	html "github.com/plainkit/html"
)

// TableColumnsSplit creates a Table Columns Split Lucide icon.
func TableColumnsSplit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-columns-split", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 14v2")),
		html.SvgPath(html.AD("M14 20v2")),
		html.SvgPath(html.AD("M14 2v2")),
		html.SvgPath(html.AD("M14 8v2")),
		html.SvgPath(html.AD("M2 15h8")),
		html.SvgPath(html.AD("M2 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H2")),
		html.SvgPath(html.AD("M2 9h8")),
		html.SvgPath(html.AD("M22 15h-4")),
		html.SvgPath(html.AD("M22 3h-2a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2")),
		html.SvgPath(html.AD("M22 9h-4")),
		html.SvgPath(html.AD("M5 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
