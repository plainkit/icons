package lucide

import x "github.com/plainkit/blox"

// RectangleEllipsis creates a Rectangle Ellipsis Lucide icon.
func RectangleEllipsis(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rectangle-ellipsis", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 12h.01"))),
		x.Child(x.Path(x.D("M17 12h.01"))),
		x.Child(x.Path(x.D("M7 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
