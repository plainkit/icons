package lucide

import x "github.com/bloxui/blox"

// TableRowsSplit creates a Table Rows Split Lucide icon.
func TableRowsSplit(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table-rows-split", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 10h2"))),
		x.Child(x.Path(x.D("M15 22v-8"))),
		x.Child(x.Path(x.D("M15 2v4"))),
		x.Child(x.Path(x.D("M2 10h2"))),
		x.Child(x.Path(x.D("M20 10h2"))),
		x.Child(x.Path(x.D("M3 19h18"))),
		x.Child(x.Path(x.D("M3 22v-6a2 2 135 0 1 2-2h14a2 2 45 0 1 2 2v6"))),
		x.Child(x.Path(x.D("M3 2v2a2 2 45 0 0 2 2h14a2 2 135 0 0 2-2V2"))),
		x.Child(x.Path(x.D("M8 10h2"))),
		x.Child(x.Path(x.D("M9 22v-8"))),
		x.Child(x.Path(x.D("M9 2v4"))),
	)
	return x.Svg(svgArgs...)
}
