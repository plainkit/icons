package lucide

import x "github.com/plainkit/html"

// Target creates a Target Lucide icon.
func Target(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-target", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("6"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
