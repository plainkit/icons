package lucide

import x "github.com/plainkit/blox"

// BetweenVerticalEnd creates a Between Vertical End Lucide icon.
func BetweenVerticalEnd(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-between-vertical-end", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("13"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("m9 22 3-3 3 3"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("13"), x.X("14"), x.Y("3"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
