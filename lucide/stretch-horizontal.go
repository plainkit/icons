package lucide

import x "github.com/plainkit/html"

// StretchHorizontal creates a Stretch Horizontal Lucide icon.
func StretchHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-stretch-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("6"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("6"), x.X("2"), x.Y("14"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
