package lucide

import x "github.com/bloxui/blox"

// Sheet creates a Sheet Lucide icon.
func Sheet(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sheet", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Line(x.X1("3"), x.X2("21"), x.Y1("9"), x.Y2("9"))),
		x.Child(x.Line(x.X1("3"), x.X2("21"), x.Y1("15"), x.Y2("15"))),
		x.Child(x.Line(x.X1("9"), x.X2("9"), x.Y1("9"), x.Y2("21"))),
		x.Child(x.Line(x.X1("15"), x.X2("15"), x.Y1("9"), x.Y2("21"))),
	)
	return x.Svg(svgArgs...)
}
