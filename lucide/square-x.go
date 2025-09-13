package lucide

import x "github.com/bloxui/blox"

// SquareX creates a Square X Lucide icon.
func SquareX(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("m9 9 6 6"))),
	)
	return x.Svg(svgArgs...)
}
