package lucide

import x "github.com/plainkit/html"

// CalendarCheck2 creates a Calendar Check 2 Lucide icon.
func CalendarCheck2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-check-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 14V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("m16 20 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
