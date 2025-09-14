package lucide

import x "github.com/bloxui/blox"

// PanelLeftDashed creates a Panel Left Dashed Lucide icon.
func PanelLeftDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-left-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 14v1"))),
		x.Child(x.Path(x.D("M9 19v2"))),
		x.Child(x.Path(x.D("M9 3v2"))),
		x.Child(x.Path(x.D("M9 9v1"))),
	)
	return x.Svg(svgArgs...)
}
