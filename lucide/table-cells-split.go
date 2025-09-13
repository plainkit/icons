package lucide

import x "github.com/bloxui/blox"

// TableCellsSplit creates a Table Cells Split Lucide icon.
func TableCellsSplit(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table-cells-split", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 15V9"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
