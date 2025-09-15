package lucide

import x "github.com/plainkit/blox"

// AlignStartHorizontal creates a Align Start Horizontal Lucide icon.
func AlignStartHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-start-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("16"), x.X("4"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("9"), x.X("14"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 2H2"))),
	)
	return x.Svg(svgArgs...)
}
