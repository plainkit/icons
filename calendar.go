package lucide

import x "github.com/bloxui/blox"

// Calendar creates a Calendar Lucide icon.
func Calendar(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-calendar", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 2v4"))),
		x.Child(x.Path(x.D("M16 2v4"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 10h18"))),
	)
	return x.Svg(svgArgs...)
}
