package lucide

import x "github.com/bloxui/blox"

// RectangleVertical creates a Rectangle Vertical Lucide icon.
func RectangleVertical(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rectangle-vertical", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("20"), x.X("6"), x.Y("2"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
