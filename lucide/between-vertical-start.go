package lucide

import x "github.com/plainkit/html"

// BetweenVerticalStart creates a Between Vertical Start Lucide icon.
func BetweenVerticalStart(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-between-vertical-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("13"), x.X("3"), x.Y("8"), x.Rx("1"))),
		x.Child(x.Path(x.D("m15 2-3 3-3-3"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("13"), x.X("14"), x.Y("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
