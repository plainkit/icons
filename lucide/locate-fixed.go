package lucide

import x "github.com/bloxui/blox"

// LocateFixed creates a Locate Fixed Lucide icon.
func LocateFixed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-locate-fixed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("2"), x.X2("5"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("19"), x.X2("22"), x.Y1("12"), x.Y2("12"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("2"), x.Y2("5"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("19"), x.Y2("22"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("7"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
