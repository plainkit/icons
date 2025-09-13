package lucide

import x "github.com/bloxui/blox"

// UserRoundCheck creates a User Round Check Lucide icon.
func UserRoundCheck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-user-round-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 13.292-6"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("m16 19 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
