package lucide

import x "github.com/bloxui/blox"

// AlignVerticalJustifyEnd creates a Align Vertical Justify End Lucide icon.
func AlignVerticalJustifyEnd(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-align-vertical-justify-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("6"), x.X("5"), x.Y("12"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("6"), x.X("7"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 22h20"))),
	)
	return x.Svg(svgArgs...)
}
