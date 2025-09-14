package lucide

import x "github.com/bloxui/blox"

// ChartBarStacked creates a Chart Bar Stacked Lucide icon.
func ChartBarStacked(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar-stacked", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 13v4"))),
		x.Child(x.Path(x.D("M15 5v4"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Rect(x.RectWidth("9"), x.RectHeight("4"), x.X("7"), x.Y("13"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("4"), x.X("7"), x.Y("5"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
