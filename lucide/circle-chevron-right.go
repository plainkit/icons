package lucide

import x "github.com/plainkit/html"

// CircleChevronRight creates a Circle Chevron Right Lucide icon.
func CircleChevronRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-chevron-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m10 8 4 4-4 4"))),
	)
	return x.Svg(svgArgs...)
}
