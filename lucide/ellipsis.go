package lucide

import x "github.com/plainkit/html"

// Ellipsis creates a Ellipsis Lucide icon.
func Ellipsis(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ellipsis", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
