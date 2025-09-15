package lucide

import x "github.com/plainkit/blox"

// UserRound creates a User Round Lucide icon.
func UserRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("M20 21a8 8 0 0 0-16 0"))),
	)
	return x.Svg(svgArgs...)
}
