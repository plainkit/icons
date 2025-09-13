package lucide

import x "github.com/bloxui/blox"

// ArrowDownUp creates a Arrow Down Up Lucide icon.
func ArrowDownUp(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 16 4 4 4-4"))),
		x.Child(x.Path(x.D("M7 20V4"))),
		x.Child(x.Path(x.D("m21 8-4-4-4 4"))),
		x.Child(x.Path(x.D("M17 4v16"))),
	)
	return x.Svg(svgArgs...)
}
