package lucide

import x "github.com/plainkit/blox"

// AlignHorizontalDistributeStart creates a Align Horizontal Distribute Start Lucide icon.
func AlignHorizontalDistributeStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-distribute-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("4"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("14"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M4 2v20"))),
		x.Child(x.Path(x.D("M14 2v20"))),
	)
	return x.Svg(svgArgs...)
}
