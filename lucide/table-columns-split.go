package lucide

import x "github.com/bloxui/blox"

// TableColumnsSplit creates a Table Columns Split Lucide icon.
func TableColumnsSplit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-table-columns-split", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 14v2"))),
		x.Child(x.Path(x.D("M14 20v2"))),
		x.Child(x.Path(x.D("M14 2v2"))),
		x.Child(x.Path(x.D("M14 8v2"))),
		x.Child(x.Path(x.D("M2 15h8"))),
		x.Child(x.Path(x.D("M2 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H2"))),
		x.Child(x.Path(x.D("M2 9h8"))),
		x.Child(x.Path(x.D("M22 15h-4"))),
		x.Child(x.Path(x.D("M22 3h-2a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2"))),
		x.Child(x.Path(x.D("M22 9h-4"))),
		x.Child(x.Path(x.D("M5 3v18"))),
	)
	return x.Svg(svgArgs...)
}
