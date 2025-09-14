package lucide

import x "github.com/bloxui/blox"

// SquareRadical creates a Square Radical Lucide icon.
func SquareRadical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-radical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 12h2l2 5 2-10h4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
