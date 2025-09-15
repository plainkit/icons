package lucide

import x "github.com/plainkit/blox"

// PanelBottomClose creates a Panel Bottom Close Lucide icon.
func PanelBottomClose(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-bottom-close", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("m15 8-3 3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
