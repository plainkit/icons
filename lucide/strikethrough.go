package lucide

import x "github.com/bloxui/blox"

// Strikethrough creates a Strikethrough Lucide icon.
func Strikethrough(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-strikethrough", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 4H9a3 3 0 0 0-2.83 4"))),
		x.Child(x.Path(x.D("M14 12a4 4 0 0 1 0 8H6"))),
		x.Child(x.Line(x.X1("4"), x.X2("20"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
