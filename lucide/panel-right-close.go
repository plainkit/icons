package lucide

import x "github.com/bloxui/blox"

// PanelRightClose creates a Panel Right Close Lucide icon.
func PanelRightClose(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-right-close", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 3v18"))),
		x.Child(x.Path(x.D("m8 9 3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
