package lucide

import x "github.com/bloxui/blox"

// CircleEqual creates a Circle Equal Lucide icon.
func CircleEqual(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-equal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 10h10"))),
		x.Child(x.Path(x.D("M7 14h10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
