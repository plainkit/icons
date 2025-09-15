package lucide

import x "github.com/plainkit/html"

// AlarmClock creates a Alarm Clock Lucide icon.
func AlarmClock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-alarm-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("8"))),
		x.Child(x.Path(x.D("M12 9v4l2 2"))),
		x.Child(x.Path(x.D("M5 3 2 6"))),
		x.Child(x.Path(x.D("m22 6-3-3"))),
		x.Child(x.Path(x.D("M6.38 18.7 4 21"))),
		x.Child(x.Path(x.D("M17.64 18.67 20 21"))),
	)
	return x.Svg(svgArgs...)
}
