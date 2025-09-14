package lucide

import x "github.com/bloxui/blox"

// StretchVertical creates a Stretch Vertical Lucide icon.
func StretchVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-stretch-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("20"), x.X("14"), x.Y("2"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
