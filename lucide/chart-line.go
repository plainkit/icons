package lucide

import x "github.com/plainkit/blox"

// ChartLine creates a Chart Line Lucide icon.
func ChartLine(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("m19 9-5 5-4-4-3 3"))),
	)
	return x.Svg(svgArgs...)
}
