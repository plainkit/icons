package lucide

import x "github.com/bloxui/blox"

// ChartNoAxesColumnDecreasing creates a Chart No Axes Column Decreasing Lucide icon.
func ChartNoAxesColumnDecreasing(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-no-axes-column-decreasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 21V3"))),
		x.Child(x.Path(x.D("M12 21V9"))),
		x.Child(x.Path(x.D("M19 21v-6"))),
	)
	return x.Svg(svgArgs...)
}
