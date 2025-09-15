package lucide

import x "github.com/plainkit/html"

// Transgender creates a Transgender Lucide icon.
func Transgender(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-transgender", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 16v6"))),
		x.Child(x.Path(x.D("M14 20h-4"))),
		x.Child(x.Path(x.D("M18 2h4v4"))),
		x.Child(x.Path(x.D("m2 2 7.17 7.17"))),
		x.Child(x.Path(x.D("M2 5.355V2h3.357"))),
		x.Child(x.Path(x.D("m22 2-7.17 7.17"))),
		x.Child(x.Path(x.D("M8 5 5 8"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
	)
	return x.Svg(svgArgs...)
}
