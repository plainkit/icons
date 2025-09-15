package lucide

import x "github.com/plainkit/html"

// AlignEndVertical creates a Align End Vertical Lucide icon.
func AlignEndVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-end-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("6"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("9"), x.RectHeight("6"), x.X("9"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Path(x.D("M22 22V2"))),
	)
	return x.Svg(svgArgs...)
}
