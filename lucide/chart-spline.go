package lucide

import x "github.com/bloxui/blox"

// ChartSpline creates a Chart Spline Lucide icon.
func ChartSpline(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-spline", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M7 16c.5-2 1.5-7 4-7 2 0 2 3 4 3 2.5 0 4.5-5 5-7"))),
	)
	return x.Svg(svgArgs...)
}
