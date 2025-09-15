package lucide

import x "github.com/plainkit/html"

// AlignHorizontalJustifyCenter creates a Align Horizontal Justify Center Lucide icon.
func AlignHorizontalJustifyCenter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-align-horizontal-justify-center", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("16"), x.Y("7"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 2v20"))),
	)
	return x.Svg(svgArgs...)
}
