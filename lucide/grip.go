package lucide

import x "github.com/plainkit/html"

// Grip creates a Grip Lucide icon.
func Grip(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grip", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("5"), x.R("1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("19"), x.R("1"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("19"), x.R("1"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
