package lucide

import x "github.com/bloxui/blox"

// ListX creates a List X Lucide icon.
func ListX(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-list-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M11 12H3"))),
		x.Child(x.Path(x.D("M16 19H3"))),
		x.Child(x.Path(x.D("m15.5 9.5 5 5"))),
		x.Child(x.Path(x.D("m20.5 9.5-5 5"))),
	)
	return x.Svg(svgArgs...)
}
