package lucide

import x "github.com/bloxui/blox"

// ChartPie creates a Chart Pie Lucide icon.
func ChartPie(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-pie", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 12c.552 0 1.005-.449.95-.998a10 10 0 0 0-8.953-8.951c-.55-.055-.998.398-.998.95v8a1 1 0 0 0 1 1z"))),
		x.Child(x.Path(x.D("M21.21 15.89A10 10 0 1 1 8 2.83"))),
	)
	return x.Svg(svgArgs...)
}
