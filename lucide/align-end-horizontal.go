package lucide

import x "github.com/bloxui/blox"

// AlignEndHorizontal creates a Align End Horizontal Lucide icon.
func AlignEndHorizontal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-align-end-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("16"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("9"), x.X("14"), x.Y("9"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 22H2"))),
	)
	return x.Svg(svgArgs...)
}
