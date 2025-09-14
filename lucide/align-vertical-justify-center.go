package lucide

import x "github.com/bloxui/blox"

// AlignVerticalJustifyCenter creates a Align Vertical Justify Center Lucide icon.
func AlignVerticalJustifyCenter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-justify-center", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("6"), x.X("5"), x.Y("16"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 12h20"))),
	)
	return x.Svg(svgArgs...)
}
