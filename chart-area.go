package lucide

import x "github.com/bloxui/blox"

// ChartArea creates a Chart Area Lucide icon.
func ChartArea(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-area", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Path(x.D("M7 11.207a.5.5 0 0 1 .146-.353l2-2a.5.5 0 0 1 .708 0l3.292 3.292a.5.5 0 0 0 .708 0l4.292-4.292a.5.5 0 0 1 .854.353V16a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1z"))),
	)
	return x.Svg(svgArgs...)
}
