package lucide

import x "github.com/bloxui/blox"

// PanelLeftRightDashed creates a Panel Left Right Dashed Lucide icon.
func PanelLeftRightDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-left-right-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 10V9"))),
		x.Child(x.Path(x.D("M15 15v-1"))),
		x.Child(x.Path(x.D("M15 21v-2"))),
		x.Child(x.Path(x.D("M15 5V3"))),
		x.Child(x.Path(x.D("M9 10V9"))),
		x.Child(x.Path(x.D("M9 15v-1"))),
		x.Child(x.Path(x.D("M9 21v-2"))),
		x.Child(x.Path(x.D("M9 5V3"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
