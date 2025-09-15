package lucide

import x "github.com/plainkit/blox"

// SquareScissors creates a Square Scissors Lucide icon.
func SquareScissors(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-scissors", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("20"), x.X("2"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("8"), x.R("2"))),
		x.Child(x.Path(x.D("M9.414 9.414 12 12"))),
		x.Child(x.Path(x.D("M14.8 14.8 18 18"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("m18 6-8.586 8.586"))),
	)
	return x.Svg(svgArgs...)
}
