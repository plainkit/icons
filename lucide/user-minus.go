package lucide

import x "github.com/bloxui/blox"

// UserMinus creates a User Minus Lucide icon.
func UserMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("4"))),
		x.Child(x.Line(x.X1("22"), x.X2("16"), x.Y1("11"), x.Y2("11"))),
	)
	return x.Svg(svgArgs...)
}
