package lucide

import x "github.com/plainkit/html"

// CaseLower creates a Case Lower Lucide icon.
func CaseLower(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-case-lower", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9v7"))),
		x.Child(x.Path(x.D("M14 6v10"))),
		x.Child(x.Circle(x.Cx("17.5"), x.Cy("12.5"), x.R("3.5"))),
		x.Child(x.Circle(x.Cx("6.5"), x.Cy("12.5"), x.R("3.5"))),
	)
	return x.Svg(svgArgs...)
}
