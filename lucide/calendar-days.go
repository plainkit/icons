package lucide

import x "github.com/plainkit/blox"

// CalendarDays creates a Calendar Days Lucide icon.
func CalendarDays(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-days", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 14h.01"))),
		x.Child(x.Path(x.D("M12 14h.01"))),
		x.Child(x.Path(x.D("M16 14h.01"))),
		x.Child(x.Path(x.D("M8 18h.01"))),
		x.Child(x.Path(x.D("M12 18h.01"))),
		x.Child(x.Path(x.D("M16 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
