package lucide

import x "github.com/bloxui/blox"

// ChartCandlestick creates a Chart Candlestick Lucide icon.
func ChartCandlestick(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chart-candlestick", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 5v4"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("6"), x.X("7"), x.Y("9"), x.Rx("1"))),
		x.Child(x.Path(x.D("M9 15v2"))),
		x.Child(x.Path(x.D("M17 3v2"))),
		x.Child(x.Rect(x.RectWidth("4"), x.RectHeight("8"), x.X("15"), x.Y("5"), x.Rx("1"))),
		x.Child(x.Path(x.D("M17 13v3"))),
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
	)
	return x.Svg(svgArgs...)
}
