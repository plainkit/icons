package lucide

import x "github.com/plainkit/html"

// LayoutDashboard creates a Layout Dashboard Lucide icon.
func LayoutDashboard(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-layout-dashboard", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("9"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("5"), x.X("14"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("9"), x.X("14"), x.Y("12"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("5"), x.X("3"), x.Y("16"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
