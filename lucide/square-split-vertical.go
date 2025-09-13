package lucide

import x "github.com/bloxui/blox"

// SquareSplitVertical creates a Square Split Vertical Lucide icon.
func SquareSplitVertical(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-split-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 8V5c0-1 1-2 2-2h10c1 0 2 1 2 2v3"))),
		x.Child(x.Path(x.D("M19 16v3c0 1-1 2-2 2H7c-1 0-2-1-2-2v-3"))),
		x.Child(x.Line(x.X1("4"), x.X2("20"), x.Y1("12"), x.Y2("12"))),
	)
	return x.Svg(svgArgs...)
}
