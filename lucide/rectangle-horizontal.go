package lucide

import x "github.com/bloxui/blox"

// RectangleHorizontal creates a Rectangle Horizontal Lucide icon.
func RectangleHorizontal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rectangle-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
