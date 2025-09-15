package lucide

import x "github.com/plainkit/blox"

// ChartColumnStacked creates a Chart Column Stacked Lucide icon.
func ChartColumnStacked(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-column-stacked", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 13H7"))),
		x.Child(x.Path(x.D("M19 9h-4"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("12"), x.X("15"), x.Y("5"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("9"), x.X("7"), x.Y("8"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
