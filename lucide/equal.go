package lucide

import x "github.com/bloxui/blox"

// Equal creates a Equal Lucide icon.
func Equal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-equal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("5"), x.X2("19"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("5"), x.X2("19"), x.Y1("15"), x.Y2("15"))),
	)
	return x.Svg(svgArgs...)
}
