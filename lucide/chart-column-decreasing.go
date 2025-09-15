package lucide

import x "github.com/plainkit/html"

// ChartColumnDecreasing creates a Chart Column Decreasing Lucide icon.
func ChartColumnDecreasing(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-column-decreasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 17V9"))),
		x.Child(x.Path(x.D("M18 17v-3"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M8 17V5"))),
	)
	return x.Svg(svgArgs...)
}
