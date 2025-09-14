package lucide

import x "github.com/bloxui/blox"

// CalendarX2 creates a Calendar X 2 Lucide icon.
func CalendarX2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-x-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("m17 22 5-5"))),
		x.Child(x.Path(x.D("m17 17 5 5"))),
	)
	return x.Svg(svgArgs...)
}
