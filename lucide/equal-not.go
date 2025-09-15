package lucide

import x "github.com/plainkit/html"

// EqualNot creates a Equal Not Lucide icon.
func EqualNot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-equal-not", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("5"), x.X2("19"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("5"), x.X2("19"), x.Y1("15"), x.Y2("15"))),
		x.Child(x.Line(x.X1("19"), x.X2("5"), x.Y1("5"), x.Y2("19"))),
	)
	return x.Svg(svgArgs...)
}
