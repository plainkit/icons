package lucide

import x "github.com/bloxui/blox"

// ChartBarBig creates a Chart Bar Big Lucide icon.
func ChartBarBig(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Rect(x.X("7"), x.Y("13"), x.RectWidth("9"), x.RectHeight("4"), x.Rx("1"))),
		x.Child(x.Rect(x.X("7"), x.Y("5"), x.RectWidth("12"), x.RectHeight("4"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
