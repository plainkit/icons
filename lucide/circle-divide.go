package lucide

import x "github.com/bloxui/blox"

// CircleDivide creates a Circle Divide Lucide icon.
func CircleDivide(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-divide", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("16"), x.Y2("16"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("8"), x.Y2("8"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
