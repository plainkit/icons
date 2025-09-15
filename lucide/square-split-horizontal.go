package lucide

import x "github.com/plainkit/blox"

// SquareSplitHorizontal creates a Square Split Horizontal Lucide icon.
func SquareSplitHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-split-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 19H5c-1 0-2-1-2-2V7c0-1 1-2 2-2h3"))),
		x.Child(x.Path(x.D("M16 5h3c1 0 2 1 2 2v10c0 1-1 2-2 2h-3"))),
		x.Child(x.Line(x.X1("12"), x.X2("12"), x.Y1("4"), x.Y2("20"))),
	)
	return x.Svg(svgArgs...)
}
