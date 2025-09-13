package lucide

import x "github.com/bloxui/blox"

// AlignHorizontalSpaceBetween creates a Align Horizontal Space Between Lucide icon.
func AlignHorizontalSpaceBetween(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-space-between", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("3"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("15"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 2v20"))),
		x.Child(x.Path(x.D("M21 2v20"))),
	)
	return x.Svg(svgArgs...)
}
