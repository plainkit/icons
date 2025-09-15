package lucide

import x "github.com/plainkit/html"

// LayoutTemplate creates a Layout Template Lucide icon.
func LayoutTemplate(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-layout-template", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("7"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("9"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("7"), x.X("16"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
