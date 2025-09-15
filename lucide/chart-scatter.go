package lucide

import x "github.com/plainkit/html"

// ChartScatter creates a Chart Scatter Lucide icon.
func ChartScatter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-scatter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("7.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("18.5"), x.Cy("5.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("11.5"), x.Cy("11.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("7.5"), x.Cy("16.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Circle(x.Cx("17.5"), x.Cy("14.5"), x.R(".5"), x.Fill("currentColor"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
	)
	return x.Svg(svgArgs...)
}
