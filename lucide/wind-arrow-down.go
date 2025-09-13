package lucide

import x "github.com/bloxui/blox"

// WindArrowDown creates a Wind Arrow Down Lucide icon.
func WindArrowDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-wind-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2v8"))),
		x.Child(x.Path(x.D("M12.8 21.6A2 2 0 1 0 14 18H2"))),
		x.Child(x.Path(x.D("M17.5 10a2.5 2.5 0 1 1 2 4H2"))),
		x.Child(x.Path(x.D("m6 6 4 4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
