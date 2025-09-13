package lucide

import x "github.com/bloxui/blox"

// CalendarFold creates a Calendar Fold Lucide icon.
func CalendarFold(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-fold", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 17V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h11Z"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M15 22v-4a2 2 0 0 1 2-2h4"))),
	)
	return x.Svg(svgArgs...)
}
