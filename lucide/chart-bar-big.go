package lucide

import x "github.com/plainkit/html"

// ChartBarBig creates a Chart Bar Big Lucide icon.
func ChartBarBig(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chart-bar-big", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3v16a2 2 0 0 0 2 2h16"))),
		x.Child(x.Rect(x.RectWidth("9"), x.RectHeight("4"), x.X("7"), x.Y("13"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("12"), x.RectHeight("4"), x.X("7"), x.Y("5"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
