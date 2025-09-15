package lucide

import x "github.com/plainkit/blox"

// ChartGantt creates a Chart Gantt Lucide icon.
func ChartGantt(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-gantt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 6h8"))),
		x.Child(x.Path(x.D("M12 16h6"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M8 11h7"))),
	)
	return x.Svg(svgArgs...)
}
