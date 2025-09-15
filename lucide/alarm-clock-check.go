package lucide

import x "github.com/plainkit/html"

// AlarmClockCheck creates a Alarm Clock Check Lucide icon.
func AlarmClockCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-alarm-clock-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("8"))),
		x.Child(x.Path(x.D("M5 3 2 6"))),
		x.Child(x.Path(x.D("m22 6-3-3"))),
		x.Child(x.Path(x.D("M6.38 18.7 4 21"))),
		x.Child(x.Path(x.D("M17.64 18.67 20 21"))),
		x.Child(x.Path(x.D("m9 13 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
