package lucide

import x "github.com/bloxui/blox"

// SquareBottomDashedScissors creates a Square Bottom Dashed Scissors Lucide icon.
func SquareBottomDashedScissors(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-bottom-dashed-scissors", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M10 22H8"))),
		x.Child(x.Path(x.D("M16 22h-2"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("8"), x.R("2"))),
		x.Child(x.Path(x.D("M9.414 9.414 12 12"))),
		x.Child(x.Path(x.D("M14.8 14.8 18 18"))),
		x.Child(x.Circle(x.Cx("8"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("m18 6-8.586 8.586"))),
	)
	return x.Svg(svgArgs...)
}
