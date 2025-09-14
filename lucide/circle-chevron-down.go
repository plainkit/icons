package lucide

import x "github.com/bloxui/blox"

// CircleChevronDown creates a Circle Chevron Down Lucide icon.
func CircleChevronDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-chevron-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m16 10-4 4-4-4"))),
	)
	return x.Svg(svgArgs...)
}
