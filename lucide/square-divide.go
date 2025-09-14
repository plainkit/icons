package lucide

import x "github.com/bloxui/blox"

// SquareDivide creates a Square Divide Lucide icon.
func SquareDivide(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-divide", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("16"), x.Y2("16"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("8"), x.Y2("8"))),
	)
	return x.Svg(svgArgs...)
}
