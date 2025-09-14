package lucide

import x "github.com/bloxui/blox"

// Variable creates a Variable Lucide icon.
func Variable(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-variable", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 21s-4-3-4-9 4-9 4-9"))),
		x.Child(x.Path(x.D("M16 3s4 3 4 9-4 9-4 9"))),
		x.Child(x.Line(x.X1("15"), x.X2("9"), x.Y1("9"), x.Y2("15"))),
		x.Child(x.Line(x.X1("9"), x.X2("15"), x.Y1("9"), x.Y2("15"))),
	)
	return x.Svg(svgArgs...)
}
