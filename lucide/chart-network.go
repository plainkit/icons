package lucide

import x "github.com/plainkit/html"

// ChartNetwork creates a Chart Network Lucide icon.
func ChartNetwork(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-network", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m13.11 7.664 1.78 2.672"))),
		x.Child(x.Path(x.D("m14.162 12.788-3.324 1.424"))),
		x.Child(x.Path(x.D("m20 4-6.06 1.515"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("6"), x.R("2"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("12"), x.R("2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("15"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
