package lucide

import x "github.com/plainkit/html"

// BetweenHorizontalEnd creates a Between Horizontal End Lucide icon.
func BetweenHorizontalEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-between-horizontal-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("13"), x.RectHeight("7"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("m22 15-3-3 3-3"))),
		x.Child(x.Rect(x.RectWidth("13"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
