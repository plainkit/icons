package lucide

import x "github.com/plainkit/blox"

// CalendarClock creates a Calendar Clock Lucide icon.
func CalendarClock(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-clock", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 14v2.2l1.6 1"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 7.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.5"))),
		x.Child(x.Path(x.D("M3 10h5"))),
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("16"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
