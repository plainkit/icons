package lucide

import x "github.com/bloxui/blox"

// AlignVerticalDistributeCenter creates a Align Vertical Distribute Center Lucide icon.
func AlignVerticalDistributeCenter(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-distribute-center", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 17h-3"))),
		x.Child(x.Path(x.D("M22 7h-5"))),
		x.Child(x.Path(x.D("M5 17H2"))),
		x.Child(x.Path(x.D("M7 7H2"))),
		x.Child(x.Rect(x.X("5"), x.Y("14"), x.RectWidth("14"), x.RectHeight("6"), x.Rx("2"))),
		x.Child(x.Rect(x.X("7"), x.Y("4"), x.RectWidth("10"), x.RectHeight("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
