package lucide

import x "github.com/bloxui/blox"

// Turntable creates a Turntable Lucide icon.
func Turntable(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-turntable", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12.01h.01"))),
		x.Child(x.Path(x.D("M18 8v4a8 8 0 0 1-1.07 4"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("12"), x.R("4"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
