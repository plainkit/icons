package lucide

import x "github.com/plainkit/blox"

// Locate creates a Locate Lucide icon.
func Locate(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-locate", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("2"), x.X2("5"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("19"), x.X2("22"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("2"), x.Y2("5"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("19"), x.Y2("22"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("7"))),
	)
	return x.Svg(svgArgs...)
}
