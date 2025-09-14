package lucide

import x "github.com/bloxui/blox"

// TimerOff creates a Timer Off Lucide icon.
func TimerOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-timer-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 2h4"))),
		x.Child(x.Path(x.D("M4.6 11a8 8 0 0 0 1.7 8.7 8 8 0 0 0 8.7 1.7"))),
		x.Child(x.Path(x.D("M7.4 7.4a8 8 0 0 1 10.3 1 8 8 0 0 1 .9 10.2"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M12 12v-2"))),
	)
	return x.Svg(svgArgs...)
}
