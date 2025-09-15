package lucide

import x "github.com/plainkit/blox"

// SquareUserRound creates a Square User Round Lucide icon.
func SquareUserRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-user-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 21a6 6 0 0 0-12 0"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("11"), x.R("4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
