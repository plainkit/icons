package lucide

import x "github.com/bloxui/blox"

// CircleDollarSign creates a Circle Dollar Sign Lucide icon.
func CircleDollarSign(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-dollar-sign", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"))),
		x.Child(x.Path(x.D("M12 18V6"))),
	)
	return x.Svg(svgArgs...)
}
