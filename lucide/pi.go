package lucide

import x "github.com/bloxui/blox"

// Pi creates a Pi Lucide icon.
func Pi(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pi", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("9"), x.X2("9"), x.Y1("4"), x.Y2("20"))),
		x.Child(x.Path(x.D("M4 7c0-1.7 1.3-3 3-3h13"))),
		x.Child(x.Path(x.D("M18 20c-1.7 0-3-1.3-3-3V4"))),
	)
	return x.Svg(svgArgs...)
}
