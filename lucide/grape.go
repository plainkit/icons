package lucide

import x "github.com/plainkit/blox"

// Grape creates a Grape Lucide icon.
func Grape(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-grape", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 5V2l-5.89 5.89"))),
		x.Child(x.Circle(x.Cx("16.6"), x.Cy("15.89"), x.R("3"))),
		x.Child(x.Circle(x.Cx("8.11"), x.Cy("7.4"), x.R("3"))),
		x.Child(x.Circle(x.Cx("12.35"), x.Cy("11.65"), x.R("3"))),
		x.Child(x.Circle(x.Cx("13.91"), x.Cy("5.85"), x.R("3"))),
		x.Child(x.Circle(x.Cx("18.15"), x.Cy("10.09"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6.56"), x.Cy("13.2"), x.R("3"))),
		x.Child(x.Circle(x.Cx("10.8"), x.Cy("17.44"), x.R("3"))),
		x.Child(x.Circle(x.Cx("5"), x.Cy("19"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
