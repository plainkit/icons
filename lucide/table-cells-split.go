package lucide

import (
	html "github.com/plainkit/html"
)

// TableCellsSplit creates a Table Cells Split Lucide icon.
func TableCellsSplit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-cells-split", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 15V9"))),
		html.Child(html.SvgPath(html.AD("M3 15h18"))),
		html.Child(html.SvgPath(html.AD("M3 9h18"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
