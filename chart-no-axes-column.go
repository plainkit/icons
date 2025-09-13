package lucide

import x "github.com/bloxui/blox"

// ChartNoAxesColumn creates a Chart No Axes Column Lucide icon.
func ChartNoAxesColumn(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-no-axes-column", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 21v-6"))),
		x.Child(x.Path(x.D("M12 21V3"))),
		x.Child(x.Path(x.D("M19 21V9"))),
	)
	return x.Svg(svgArgs...)
}
