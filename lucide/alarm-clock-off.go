package lucide

import x "github.com/bloxui/blox"

// AlarmClockOff creates a Alarm Clock Off Lucide icon.
func AlarmClockOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-alarm-clock-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6.87 6.87a8 8 0 1 0 11.26 11.26"))),
		x.Child(x.Path(x.D("M19.9 14.25a8 8 0 0 0-9.15-9.15"))),
		x.Child(x.Path(x.D("m22 6-3-3"))),
		x.Child(x.Path(x.D("M6.26 18.67 4 21"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M4 4 2 6"))),
	)
	return x.Svg(svgArgs...)
}
