package lucide

import x "github.com/plainkit/html"

// ChartNoAxesCombined creates a Chart No Axes Combined Lucide icon.
func ChartNoAxesCombined(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-no-axes-combined", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 16v5"))),
		x.Child(x.Path(x.D("M16 14v7"))),
		x.Child(x.Path(x.D("M20 10v11"))),
		x.Child(x.Path(x.D("m22 3-8.646 8.646a.5.5 0 0 1-.708 0L9.354 8.354a.5.5 0 0 0-.707 0L2 15"))),
		x.Child(x.Path(x.D("M4 18v3"))),
		x.Child(x.Path(x.D("M8 14v7"))),
	)
	return x.Svg(svgArgs...)
}
