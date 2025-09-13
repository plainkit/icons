package lucide

import x "github.com/bloxui/blox"

// UserRoundPen creates a User Round Pen Lucide icon.
func UserRoundPen(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-user-round-pen", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 21a8 8 0 0 1 10.821-7.487"))),
		x.Child(x.Path(x.D("M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("8"), x.R("5"))),
	)
	return x.Svg(svgArgs...)
}
