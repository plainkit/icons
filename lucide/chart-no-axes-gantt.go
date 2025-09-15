package lucide

import x "github.com/plainkit/html"

// ChartNoAxesGantt creates a Chart No Axes Gantt Lucide icon.
func ChartNoAxesGantt(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-no-axes-gantt", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 5h12"))),
		x.Child(x.Path(x.D("M4 12h10"))),
		x.Child(x.Path(x.D("M12 19h8"))),
	)
	return x.Svg(svgArgs...)
}
