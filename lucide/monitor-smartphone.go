package lucide

import x "github.com/plainkit/html"

// MonitorSmartphone creates a Monitor Smartphone Lucide icon.
func MonitorSmartphone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-smartphone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 8V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8"))),
		x.Child(x.Path(x.D("M10 19v-3.96 3.15"))),
		x.Child(x.Path(x.D("M7 19h5"))),
		x.Child(x.Rect(x.RectWidth("6"), x.RectHeight("10"), x.X("16"), x.Y("12"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
