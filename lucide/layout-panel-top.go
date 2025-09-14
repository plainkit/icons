package lucide

import x "github.com/bloxui/blox"

// LayoutPanelTop creates a Layout Panel Top Lucide icon.
func LayoutPanelTop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-layout-panel-top", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("7"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("14"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
