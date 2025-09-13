package lucide

import x "github.com/bloxui/blox"

// Frown creates a Frown Lucide icon.
func Frown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-frown", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M16 16s-1.5-2-4-2-4 2-4 2"))),
		x.Child(x.Line(x.X1("9"), x.X2("9.01"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("9"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
