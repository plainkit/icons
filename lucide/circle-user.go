package lucide

import x "github.com/plainkit/html"

// CircleUser creates a Circle User Lucide icon.
func CircleUser(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-user", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("M7 20.662V19a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.662"))),
	)
	return x.Svg(svgArgs...)
}
