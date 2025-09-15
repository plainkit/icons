package lucide

import x "github.com/plainkit/blox"

// CalendarOff creates a Calendar Off Lucide icon.
func CalendarOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4.2 4.2A2 2 0 0 0 3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 1.82-1.18"))),
		x.Child(x.Path(x.D("M21 15.5V6a2 2 0 0 0-2-2H9.5"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M3 10h7"))),
		x.Child(x.Path(x.D("M21 10h-5.5"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
