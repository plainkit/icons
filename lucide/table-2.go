package lucide

import x "github.com/bloxui/blox"

// Table2 creates a Table 2 Lucide icon.
func Table2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-table-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 3H5a2 2 0 0 0-2 2v4m6-6h10a2 2 0 0 1 2 2v4M9 3v18m0 0h10a2 2 0 0 0 2-2V9M9 21H5a2 2 0 0 1-2-2V9m0 0h18"))),
	)
	return x.Svg(svgArgs...)
}
