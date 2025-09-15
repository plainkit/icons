package lucide

import x "github.com/plainkit/blox"

// SquareMinus creates a Square Minus Lucide icon.
func SquareMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
