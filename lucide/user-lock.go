package lucide

import x "github.com/plainkit/blox"

// UserLock creates a User Lock Lucide icon.
func UserLock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-user-lock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("10"), x.Cy("7"), x.R("4"))),
		x.Child(x.Path(x.D("M10.3 15H7a4 4 0 0 0-4 4v2"))),
		x.Child(x.Path(x.D("M15 15.5V14a2 2 0 0 1 4 0v1.5"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("5"), x.X("13"), x.Y("16"), x.Rx(".899"))),
	)
	return x.Svg(svgArgs...)
}
