package lucide

import x "github.com/plainkit/html"

// ClosedCaption creates a Closed Caption Lucide icon.
func ClosedCaption(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-closed-caption", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9.17a3 3 0 1 0 0 5.66"))),
		x.Child(x.Path(x.D("M17 9.17a3 3 0 1 0 0 5.66"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("5"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
