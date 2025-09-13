package lucide

import x "github.com/bloxui/blox"

// CalendarArrowDown creates a Calendar Arrow Down Lucide icon.
func CalendarArrowDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 18 4 4 4-4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M18 14v8"))),
		x.Child(x.Path(x.D("M21 11.354V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.343"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
