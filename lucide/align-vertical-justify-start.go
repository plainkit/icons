package lucide

import x "github.com/bloxui/blox"

// AlignVerticalJustifyStart creates a Align Vertical Justify Start Lucide icon.
func AlignVerticalJustifyStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-justify-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("6"), x.X("5"), x.Y("16"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 2h20"))),
	)
	return x.Svg(svgArgs...)
}
