package lucide

import x "github.com/plainkit/html"

// LayoutGrid creates a Layout Grid Lucide icon.
func LayoutGrid(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-layout-grid", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("14"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("14"), x.Y("14"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
