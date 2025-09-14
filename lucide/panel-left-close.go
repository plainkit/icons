package lucide

import x "github.com/bloxui/blox"

// PanelLeftClose creates a Panel Left Close Lucide icon.
func PanelLeftClose(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-left-close", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Path(x.D("m16 15-3-3 3-3"))),
	)
	return x.Svg(svgArgs...)
}
