package lucide

import x "github.com/bloxui/blox"

// CalendarMinus2 creates a Calendar Minus 2 Lucide icon.
func CalendarMinus2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-minus-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M10 16h4"))),
	)
	return x.Svg(svgArgs...)
}
