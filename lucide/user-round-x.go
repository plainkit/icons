package lucide

import x "github.com/plainkit/html"

// UserRoundX creates a User Round X Lucide icon.
func UserRoundX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-round-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 11.873-7"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("m17 17 5 5"))),
		x.Child(x.Path(x.D("m22 17-5 5"))),
	)
	return x.Svg(svgArgs...)
}
