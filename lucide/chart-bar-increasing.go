package lucide

import x "github.com/plainkit/html"

// ChartBarIncreasing creates a Chart Bar Increasing Lucide icon.
func ChartBarIncreasing(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar-increasing", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M7 11h8"))),
		x.Child(x.Path(x.D("M7 16h12"))),
		x.Child(x.Path(x.D("M7 6h3"))),
	)
	return x.Svg(svgArgs...)
}
