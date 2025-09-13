package lucide

import x "github.com/bloxui/blox"

// MonitorOff creates a Monitor Off Lucide icon.
func MonitorOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-monitor-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 17H4a2 2 0 0 1-2-2V5c0-1.5 1-2 1-2"))),
		x.Child(x.Path(x.D("M22 15V5a2 2 0 0 0-2-2H9"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
