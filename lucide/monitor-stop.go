package lucide

import x "github.com/bloxui/blox"

// MonitorStop creates a Monitor Stop Lucide icon.
func MonitorStop(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-monitor-stop", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Rect(x.X("2"), x.Y("3"), x.RectWidth("20"), x.RectHeight("14"), x.Rx("2"))),
		x.Child(x.Rect(x.X("9"), x.Y("7"), x.RectWidth("6"), x.RectHeight("6"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
