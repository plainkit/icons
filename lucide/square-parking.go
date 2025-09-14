package lucide

import x "github.com/bloxui/blox"

// SquareParking creates a Square Parking Lucide icon.
func SquareParking(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-parking", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 17V7h4a3 3 0 0 1 0 6H9"))),
	)
	return x.Svg(svgArgs...)
}
