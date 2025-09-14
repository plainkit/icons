package lucide

import x "github.com/bloxui/blox"

// AlignHorizontalJustifyStart creates a Align Horizontal Justify Start Lucide icon.
func AlignHorizontalJustifyStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-justify-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("6"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("16"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 2v20"))),
	)
	return x.Svg(svgArgs...)
}
