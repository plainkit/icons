package lucide

import x "github.com/plainkit/html"

// MonitorDot creates a Monitor Dot Lucide icon.
func MonitorDot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M22 12.307V15a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8.693"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Circle(x.Cx("19"), x.Cy("6"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
