package lucide

import x "github.com/bloxui/blox"

// PanelRightOpen creates a Panel Right Open Lucide icon.
func PanelRightOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-right-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 3v18"))),
		x.Child(x.Path(x.D("m10 15-3-3 3-3"))),
	)
	return x.Svg(svgArgs...)
}
