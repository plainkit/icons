package lucide

import x "github.com/bloxui/blox"

// SquareFunction creates a Square Function Lucide icon.
func SquareFunction(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-square-function", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"), x.Ry("2"))),
		x.Child(x.Path(x.D("M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3"))),
		x.Child(x.Path(x.D("M9 11.2h5.7"))),
	)
	return x.Svg(svgArgs...)
}
