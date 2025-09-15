package lucide

import x "github.com/plainkit/blox"

// AlignVerticalDistributeStart creates a Align Vertical Distribute Start Lucide icon.
func AlignVerticalDistributeStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-distribute-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("6"), x.X("5"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 14h20"))),
		x.Child(x.Path(x.D("M2 4h20"))),
	)
	return x.Svg(svgArgs...)
}
