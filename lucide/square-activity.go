package lucide

import x "github.com/plainkit/html"

// SquareActivity creates a Square Activity Lucide icon.
func SquareActivity(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-activity", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M17 12h-2l-2 5-2-10-2 5H7"))),
	)
	return x.Svg(svgArgs...)
}
