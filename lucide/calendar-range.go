package lucide

import x "github.com/bloxui/blox"

// CalendarRange creates a Calendar Range Lucide icon.
func CalendarRange(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-range", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M17 14h-6"))),
		x.Child(x.Path(x.D("M13 18H7"))),
		x.Child(x.Path(x.D("M7 14h.01"))),
		x.Child(x.Path(x.D("M17 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
