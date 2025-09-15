package lucide

import x "github.com/plainkit/html"

// CirclePercent creates a Circle Percent Lucide icon.
func CirclePercent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Path(x.D("M15 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
