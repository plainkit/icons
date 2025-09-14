package lucide

import x "github.com/bloxui/blox"

// CalendarX creates a Calendar X Lucide icon.
func CalendarX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("m14 14-4 4"))),
		x.Child(x.Path(x.D("m10 14 4 4"))),
	)
	return x.Svg(svgArgs...)
}
