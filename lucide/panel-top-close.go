package lucide

import x "github.com/plainkit/blox"

// PanelTopClose creates a Panel Top Close Lucide icon.
func PanelTopClose(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-top-close", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("m9 16 3-3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
