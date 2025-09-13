package lucide

import x "github.com/bloxui/blox"

// BetweenHorizontalStart creates a Between Horizontal Start Lucide icon.
func BetweenHorizontalStart(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-between-horizontal-start", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("13"), x.RectHeight("7"), x.X("8"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Path(x.D("m2 9 3 3-3 3"))),
		x.Child(x.Rect(x.RectWidth("13"), x.RectHeight("7"), x.X("8"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
