package lucide

import x "github.com/plainkit/html"

// PillBottle creates a Pill Bottle Lucide icon.
func PillBottle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pill-bottle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 11h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4"))),
		x.Child(x.Path(x.D("M6 7v13a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("5"), x.X("4"), x.Y("2"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
