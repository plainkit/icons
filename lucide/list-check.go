package lucide

import x "github.com/bloxui/blox"

// ListCheck creates a List Check Lucide icon.
func ListCheck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M16 12H3"))),
		x.Child(x.Path(x.D("M11 19H3"))),
		x.Child(x.Path(x.D("m15 18 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
