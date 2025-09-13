package lucide

import x "github.com/bloxui/blox"

// CircleMinus creates a Circle Minus Lucide icon.
func CircleMinus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
