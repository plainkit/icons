package lucide

import x "github.com/plainkit/blox"

// Italic creates a Italic Lucide icon.
func Italic(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-italic", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("19"), x.X2("10"), x.Y1("4"), x.Y2("4"))),
		x.Child(x.Line(x.X1("14"), x.X2("5"), x.Y1("20"), x.Y2("20"))),
		x.Child(x.Line(x.X1("15"), x.X2("9"), x.Y1("4"), x.Y2("20"))),
	)
	return x.Svg(svgArgs...)
}
