package lucide

import x "github.com/bloxui/blox"

// AlignHorizontalDistributeCenter creates a Align Horizontal Distribute Center Lucide icon.
func AlignHorizontalDistributeCenter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-distribute-center", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("4"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("14"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M17 22v-5"))),
		x.Child(x.Path(x.D("M17 7V2"))),
		x.Child(x.Path(x.D("M7 22v-3"))),
		x.Child(x.Path(x.D("M7 5V2"))),
	)
	return x.Svg(svgArgs...)
}
