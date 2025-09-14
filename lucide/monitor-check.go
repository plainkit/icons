package lucide

import x "github.com/bloxui/blox"

// MonitorCheck creates a Monitor Check Lucide icon.
func MonitorCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 10 2 2 4-4"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
	)
	return x.Svg(svgArgs...)
}
