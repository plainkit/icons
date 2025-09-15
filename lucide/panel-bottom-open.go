package lucide

import x "github.com/plainkit/blox"

// PanelBottomOpen creates a Panel Bottom Open Lucide icon.
func PanelBottomOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-bottom-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 15h18"))),
		x.Child(x.Path(x.D("m9 10 3-3 3 3"))),
	)
	return x.Svg(svgArgs...)
}
