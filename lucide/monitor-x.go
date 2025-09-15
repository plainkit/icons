package lucide

import x "github.com/plainkit/html"

// MonitorX creates a Monitor X Lucide icon.
func MonitorX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14.5 12.5-5-5"))),
		x.Child(x.Path(x.D("m9.5 12.5 5-5"))),
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("14"), x.X("2"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
	)
	return x.Svg(svgArgs...)
}
