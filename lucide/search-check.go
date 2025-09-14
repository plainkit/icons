package lucide

import x "github.com/bloxui/blox"

// SearchCheck creates a Search Check Lucide icon.
func SearchCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-search-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m8 11 2 2 4-4"))),
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
		x.Child(x.Path(x.D("m21 21-4.3-4.3"))),
	)
	return x.Svg(svgArgs...)
}
