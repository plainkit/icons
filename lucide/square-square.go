package lucide

import x "github.com/bloxui/blox"

// SquareSquare creates a Square Square Lucide icon.
func SquareSquare(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-square", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("8"), x.X("8"), x.Y("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
