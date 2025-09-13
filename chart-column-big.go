package lucide

import x "github.com/bloxui/blox"

// ChartColumnBig creates a Chart Column Big Lucide icon.
func ChartColumnBig(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-column-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Rect(x.X("15"), x.Y("5"), x.RectWidth("4"), x.RectHeight("12"), x.Rx("1"))),
		x.Child(x.Rect(x.X("7"), x.Y("8"), x.RectWidth("4"), x.RectHeight("9"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
