package lucide

import x "github.com/bloxui/blox"

// Frame creates a Frame Lucide icon.
func Frame(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-frame", args)
	svgArgs = append(svgArgs,
		x.Child(x.Line(x.X1("22"), x.X2("2"), x.Y1("6"), x.Y2("6"))),
		x.Child(x.Line(x.X1("22"), x.X2("2"), x.Y1("18"), x.Y2("18"))),
		x.Child(x.Line(x.X1("6"), x.X2("6"), x.Y1("2"), x.Y2("22"))),
		x.Child(x.Line(x.X1("18"), x.X2("18"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
