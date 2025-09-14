package lucide

import x "github.com/bloxui/blox"

// MonitorSpeaker creates a Monitor Speaker Lucide icon.
func MonitorSpeaker(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-monitor-speaker", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5.5 20H8"))),
		x.Child(x.Path(x.D("M17 9h.01"))),
		x.Child(x.Rect(x.RectWidth("10"), x.RectHeight("16"), x.X("12"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M8 6H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("15"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
