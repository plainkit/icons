package lucide

import x "github.com/bloxui/blox"

// Laugh creates a Laugh Lucide icon.
func Laugh(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-laugh", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M18 13a6 6 0 0 1-6 5 6 6 0 0 1-6-5h12Z"))),
		x.Child(x.Line(x.X1("9"), x.X2("9.01"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("15"), x.X2("15.01"), x.Y1("9"), x.Y2("9"))),
	)
	return x.Svg(svgArgs...)
}
