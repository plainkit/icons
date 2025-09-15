package lucide

import x "github.com/plainkit/html"

// CalendarArrowUp creates a Calendar Arrow Up Lucide icon.
func CalendarArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 18 4-4 4 4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M18 22v-8"))),
		x.Child(x.Path(x.D("M21 11.343V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h9"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
