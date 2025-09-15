package lucide

import x "github.com/plainkit/blox"

// Scissors creates a Scissors Lucide icon.
func Scissors(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scissors", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M8.12 8.12 12 12"))),
		x.Child(x.Path(x.D("M20 4 8.12 15.88"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("M14.8 14.8 20 20"))),
	)
	return x.Svg(svgArgs...)
}
