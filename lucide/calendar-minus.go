package lucide

import x "github.com/bloxui/blox"

// CalendarMinus creates a Calendar Minus Lucide icon.
func CalendarMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 19h6"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 15V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8.5"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
