package lucide

import x "github.com/plainkit/html"

// CircleUserRound creates a Circle User Round Lucide icon.
func CircleUserRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-user-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 20a6 6 0 0 0-12 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("4"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
