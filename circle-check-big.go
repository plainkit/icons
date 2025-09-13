package lucide

import x "github.com/bloxui/blox"

// CircleCheckBig creates a Circle Check Big Lucide icon.
func CircleCheckBig(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-check-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21.801 10A10 10 0 1 1 17 3.335"))),
		x.Child(x.Path(x.D("m9 11 3 3L22 4"))),
	)
	return x.Svg(svgArgs...)
}
