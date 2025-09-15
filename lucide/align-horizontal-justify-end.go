package lucide

import x "github.com/plainkit/blox"

// AlignHorizontalJustifyEnd creates a Align Horizontal Justify End Lucide icon.
func AlignHorizontalJustifyEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-justify-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("12"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 2v20"))),
	)
	return x.Svg(svgArgs...)
}
