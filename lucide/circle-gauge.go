package lucide

import x "github.com/plainkit/blox"

// CircleGauge creates a Circle Gauge Lucide icon.
func CircleGauge(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-gauge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15.6 2.7a10 10 0 1 0 5.7 5.7"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("M13.4 10.6 19 5"))),
	)
	return x.Svg(svgArgs...)
}
