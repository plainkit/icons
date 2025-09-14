package lucide

import x "github.com/bloxui/blox"

// CircleChevronUp creates a Circle Chevron Up Lucide icon.
func CircleChevronUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-chevron-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m8 14 4-4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
