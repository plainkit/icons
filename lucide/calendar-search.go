package lucide

import x "github.com/bloxui/blox"

// CalendarSearch creates a Calendar Search Lucide icon.
func CalendarSearch(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("M21 11.75V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.25"))),
		x.Child(x.Path(x.D("m22 22-1.875-1.875"))),
		x.Child(x.Path(x.D("M3 10h18"))),
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
