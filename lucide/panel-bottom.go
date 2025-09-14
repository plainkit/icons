package lucide

import x "github.com/bloxui/blox"

// PanelBottom creates a Panel Bottom Lucide icon.
func PanelBottom(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-bottom", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 15h18"))),
	)
	return x.Svg(svgArgs...)
}
