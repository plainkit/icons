package lucide

import x "github.com/bloxui/blox"

// AlarmSmoke creates a Alarm Smoke Lucide icon.
func AlarmSmoke(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-alarm-smoke", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 21c0-2.5 2-2.5 2-5"))),
		x.Child(x.Path(x.D("M16 21c0-2.5 2-2.5 2-5"))),
		x.Child(x.Path(x.D("m19 8-.8 3a1.25 1.25 0 0 1-1.2 1H7a1.25 1.25 0 0 1-1.2-1L5 8"))),
		x.Child(x.Path(x.D("M21 3a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a1 1 0 0 1 1-1z"))),
		x.Child(x.Path(x.D("M6 21c0-2.5 2-2.5 2-5"))),
	)
	return x.Svg(svgArgs...)
}
