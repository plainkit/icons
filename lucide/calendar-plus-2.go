package lucide

import x "github.com/bloxui/blox"

// CalendarPlus2 creates a Calendar Plus 2 Lucide icon.
func CalendarPlus2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-plus-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M10 16h4"))),
		x.Child(x.Path(x.D("M12 14v4"))),
	)
	return x.Svg(svgArgs...)
}
