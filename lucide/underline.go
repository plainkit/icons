package lucide

import x "github.com/bloxui/blox"

// Underline creates a Underline Lucide icon.
func Underline(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-underline", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 4v6a6 6 0 0 0 12 0V4"))),
		x.Child(x.Line(x.X1("4"), x.X2("20"), x.Y1("20"), x.Y2("20"))),
	)
	return x.Svg(svgArgs...)
}
