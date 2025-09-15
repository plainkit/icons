package lucide

import x "github.com/plainkit/blox"

// UsersRound creates a Users Round Lucide icon.
func UsersRound(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-users-round", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 21a8 8 0 0 0-16 0"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
		x.Child(x.Path(x.D("M22 20c0-3.37-2-6.5-4-8a5 5 0 0 0-.45-8.3"))),
	)
	return x.Svg(svgArgs...)
}
