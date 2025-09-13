package lucide

import x "github.com/bloxui/blox"

// ChartNoAxesColumnIncreasing creates a Chart No Axes Column Increasing Lucide icon.
func ChartNoAxesColumnIncreasing(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-no-axes-column-increasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 21v-6"))),
		x.Child(x.Path(x.D("M12 21V9"))),
		x.Child(x.Path(x.D("M19 21V3"))),
	)
	return x.Svg(svgArgs...)
}
