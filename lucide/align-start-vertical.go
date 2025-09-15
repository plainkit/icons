package lucide

import x "github.com/plainkit/html"

// AlignStartVertical creates a Align Start Vertical Lucide icon.
func AlignStartVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-start-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("9"), x.RectHeight("6"), x.X("6"), x.Y("14"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("6"), x.X("6"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M2 2v20"))),
	)
	return x.Svg(svgArgs...)
}
