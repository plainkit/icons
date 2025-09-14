package lucide

import x "github.com/bloxui/blox"

// SearchCode creates a Search Code Lucide icon.
func SearchCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-search-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13 13.5 2-2.5-2-2.5"))),
		x.Child(x.Path(x.D("m21 21-4.3-4.3"))),
		x.Child(x.Path(x.D("M9 8.5 7 11l2 2.5"))),
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
	)
	return x.Svg(svgArgs...)
}
