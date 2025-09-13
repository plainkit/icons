package lucide

import x "github.com/bloxui/blox"

// UserRoundSearch creates a User Round Search Lucide icon.
func UserRoundSearch(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-user-round-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 10.434-7.62"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Path(x.D("m22 22-1.9-1.9"))),
	)
	return x.Svg(svgArgs...)
}
