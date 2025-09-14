package lucide

import x "github.com/bloxui/blox"

// CalendarSync creates a Calendar Sync Lucide icon.
func CalendarSync(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calendar-sync", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 10v4h4"))),
		x.Child(x.Path(x.D("m11 14 1.535-1.605a5 5 0 0 1 8 1.5"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Path(x.D("m21 18-1.535 1.605a5 5 0 0 1-8-1.5"))),
		x.Child(x.Path(x.D("M21 22v-4h-4"))),
		x.Child(x.Path(x.D("M21 8.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h4.3"))),
		x.Child(x.Path(x.D("M3 10h4"))),
		x.Child(x.Path(x.D("M8 2v4"))),
	)
	return x.Svg(svgArgs...)
}
