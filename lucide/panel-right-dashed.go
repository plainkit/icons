package lucide

import x "github.com/plainkit/blox"

// PanelRightDashed creates a Panel Right Dashed Lucide icon.
func PanelRightDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-right-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 14v1"))),
		x.Child(x.Path(x.D("M15 19v2"))),
		x.Child(x.Path(x.D("M15 3v2"))),
		x.Child(x.Path(x.D("M15 9v1"))),
	)
	return x.Svg(svgArgs...)
}
