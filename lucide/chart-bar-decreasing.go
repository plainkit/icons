package lucide

import x "github.com/plainkit/html"

// ChartBarDecreasing creates a Chart Bar Decreasing Lucide icon.
func ChartBarDecreasing(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar-decreasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M7 11h8"))),
		x.Child(x.Path(x.D("M7 16h3"))),
		x.Child(x.Path(x.D("M7 6h12"))),
	)
	return x.Svg(svgArgs...)
}
