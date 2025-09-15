package lucide

import x "github.com/plainkit/html"

// CalendarPlus creates a Calendar Plus Lucide icon.
func CalendarPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 19h6"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M19 16v6"))),
		x.Child(x.Path(x.D("M21 12.598V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8.5"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
