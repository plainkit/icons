package lucide

import x "github.com/plainkit/blox"

// SquareStop creates a Square Stop Lucide icon.
func SquareStop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-stop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("6"), x.X("9"), x.Y("9"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
