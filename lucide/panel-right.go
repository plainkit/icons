package lucide

import x "github.com/bloxui/blox"

// PanelRight creates a Panel Right Lucide icon.
func PanelRight(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-panel-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 3v18"))),
	)
	return x.Svg(svgArgs...)
}
