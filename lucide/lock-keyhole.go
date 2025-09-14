package lucide

import x "github.com/bloxui/blox"

// LockKeyhole creates a Lock Keyhole Lucide icon.
func LockKeyhole(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lock-keyhole", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("16"), x.R("1"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("12"), x.X("3"), x.Y("10"), x.Rx("2"))),
		x.Child(x.Path(x.D("M7 10V7a5 5 0 0 1 10 0v3"))),
	)
	return x.Svg(svgArgs...)
}
