package lucide

import x "github.com/plainkit/html"

// PersonStanding creates a Person Standing Lucide icon.
func PersonStanding(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-person-standing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("1"))),
		x.Child(x.Path(x.D("m9 20 3-6 3 6"))),
		x.Child(x.Path(x.D("m6 8 6 2 6-2"))),
		x.Child(x.Path(x.D("M12 10v4"))),
	)
	return x.Svg(svgArgs...)
}
