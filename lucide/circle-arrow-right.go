package lucide

import x "github.com/plainkit/blox"

// CircleArrowRight creates a Circle Arrow Right Lucide icon.
func CircleArrowRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m12 16 4-4-4-4"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
