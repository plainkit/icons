package lucide

import x "github.com/plainkit/blox"

// ListFilterPlus creates a List Filter Plus Lucide icon.
func ListFilterPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-filter-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5H2"))),
		x.Child(x.Path(x.D("M6 12h12"))),
		x.Child(x.Path(x.D("M9 19h6"))),
		x.Child(x.Path(x.D("M16 5h6"))),
		x.Child(x.Path(x.D("M19 8V2"))),
	)
	return x.Svg(svgArgs...)
}
