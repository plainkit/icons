package lucide

import x "github.com/bloxui/blox"

// CircleParking creates a Circle Parking Lucide icon.
func CircleParking(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-parking", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M9 17V7h4a3 3 0 0 1 0 6H9"))),
	)
	return x.Svg(svgArgs...)
}
