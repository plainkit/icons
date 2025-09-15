package lucide

import x "github.com/plainkit/html"

// DollarSign creates a Dollar Sign Lucide icon.
func DollarSign(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dollar-sign", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("2"), x.Y2("22"))),
		x.Child(x.Path(x.D("M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"))),
	)
	return x.Svg(svgArgs...)
}
