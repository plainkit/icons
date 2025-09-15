package lucide

import x "github.com/plainkit/blox"

// SquarePercent creates a Square Percent Lucide icon.
func SquarePercent(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-percent", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("m15 9-6 6"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Path(x.D("M15 15h.01"))),
	)
	return x.Svg(svgArgs...)
}
