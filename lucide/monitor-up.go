package lucide

import x "github.com/plainkit/html"

// MonitorUp creates a Monitor Up Lucide icon.
func MonitorUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 10 3-3 3 3"))),
		x.Child(x.Path(x.D("M12 13V7"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
	)
	return x.Svg(svgArgs...)
}
