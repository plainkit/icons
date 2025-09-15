package lucide

import x "github.com/plainkit/blox"

// SquareChartGantt creates a Square Chart Gantt Lucide icon.
func SquareChartGantt(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-chart-gantt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 8h7"))),
		x.Child(x.Path(x.D("M8 12h6"))),
		x.Child(x.Path(x.D("M11 16h5"))),
	)
	return x.Svg(svgArgs...)
}
