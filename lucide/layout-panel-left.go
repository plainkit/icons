package lucide

import x "github.com/bloxui/blox"

// LayoutPanelLeft creates a Layout Panel Left Lucide icon.
func LayoutPanelLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-layout-panel-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("14"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("14"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
