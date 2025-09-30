package lucide

import (
	html "github.com/plainkit/html"
)

// TableCellsMerge creates a Table Cells Merge Lucide icon.
func TableCellsMerge(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-cells-merge", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 21v-6")),
		html.SvgPath(html.AD("M12 9V3")),
		html.SvgPath(html.AD("M3 15h18")),
		html.SvgPath(html.AD("M3 9h18")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
