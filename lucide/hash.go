package lucide

import x "github.com/bloxui/blox"

// Hash creates a Hash Lucide icon.
func Hash(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("4"), x.X2("20"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("4"), x.X2("20"), x.Y1("15"), x.Y2("15"))),
		x.Child(x.Line(x.X1("10"), x.X2("8"), x.Y1("3"), x.Y2("21"))),
		x.Child(x.Line(x.X1("16"), x.X2("14"), x.Y1("3"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
