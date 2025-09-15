package lucide

import x "github.com/plainkit/html"

// CircleChevronLeft creates a Circle Chevron Left Lucide icon.
func CircleChevronLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-chevron-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m14 16-4-4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
