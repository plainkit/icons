package lucide

import x "github.com/plainkit/html"

// AlarmClockPlus creates a Alarm Clock Plus Lucide icon.
func AlarmClockPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-alarm-clock-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("8"))),
		x.Child(x.Path(x.D("M5 3 2 6"))),
		x.Child(x.Path(x.D("m22 6-3-3"))),
		x.Child(x.Path(x.D("M6.38 18.7 4 21"))),
		x.Child(x.Path(x.D("M17.64 18.67 20 21"))),
		x.Child(x.Path(x.D("M12 10v6"))),
		x.Child(x.Path(x.D("M9 13h6"))),
	)
	return x.Svg(svgArgs...)
}
