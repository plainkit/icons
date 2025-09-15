package lucide

import x "github.com/plainkit/html"

// ChartColumnIncreasing creates a Chart Column Increasing Lucide icon.
func ChartColumnIncreasing(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-column-increasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 17V9"))),
		x.Child(x.Path(x.D("M18 17V5"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M8 17v-3"))),
	)
	return x.Svg(svgArgs...)
}
