package lucide

import x "github.com/plainkit/html"

// MonitorPause creates a Monitor Pause Lucide icon.
func MonitorPause(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-pause", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 13V7"))),
		x.Child(x.Path(x.D("M14 13V7"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
	)
	return x.Svg(svgArgs...)
}
