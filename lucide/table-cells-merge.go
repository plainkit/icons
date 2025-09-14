package lucide

import x "github.com/bloxui/blox"

// TableCellsMerge creates a Table Cells Merge Lucide icon.
func TableCellsMerge(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-table-cells-merge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 21v-6"))),
		x.Child(x.Path(x.D("M12 9V3"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
