package lucide

import x "github.com/plainkit/html"

// ChartBar creates a Chart Bar Lucide icon.
func ChartBar(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M7 16h8"))),
		x.Child(x.Path(x.D("M7 11h12"))),
		x.Child(x.Path(x.D("M7 6h3"))),
	)
	return x.Svg(svgArgs...)
}
