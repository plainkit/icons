package lucide

import x "github.com/bloxui/blox"

// CircleX creates a Circle X Lucide icon.
func CircleX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("m9 9 6 6"))),
	)
	return x.Svg(svgArgs...)
}
