package lucide

import x "github.com/bloxui/blox"

// SquareArrowDownLeft creates a Square Arrow Down Left Lucide icon.
func SquareArrowDownLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-down-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m16 8-8 8"))),
		x.Child(x.Path(x.D("M16 16H8V8"))),
	)
	return x.Svg(svgArgs...)
}
