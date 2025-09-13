package lucide

import x "github.com/bloxui/blox"

// UserRoundMinus creates a User Round Minus Lucide icon.
func UserRoundMinus(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-user-round-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 13.292-6"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("M22 19h-6"))),
	)
	return x.Svg(svgArgs...)
}
