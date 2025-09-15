package lucide

import x "github.com/plainkit/blox"

// CirclePoundSterling creates a Circle Pound Sterling Lucide icon.
func CirclePoundSterling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-pound-sterling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 16V9.5a1 1 0 0 1 5 0"))),
		x.Child(x.Path(x.D("M8 12h4"))),
		x.Child(x.Path(x.D("M8 16h7"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
