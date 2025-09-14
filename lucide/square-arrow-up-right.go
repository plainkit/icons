package lucide

import x "github.com/bloxui/blox"

// SquareArrowUpRight creates a Square Arrow Up Right Lucide icon.
func SquareArrowUpRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 8h8v8"))),
		x.Child(x.Path(x.D("m8 16 8-8"))),
	)
	return x.Svg(svgArgs...)
}
