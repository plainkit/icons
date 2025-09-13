package lucide

import x "github.com/bloxui/blox"

// PanelBottomDashed creates a Panel Bottom Dashed Lucide icon.
func PanelBottomDashed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-panel-bottom-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M14 15h1"))),
		x.Child(x.Path(x.D("M19 15h2"))),
		x.Child(x.Path(x.D("M3 15h2"))),
		x.Child(x.Path(x.D("M9 15h1"))),
	)
	return x.Svg(svgArgs...)
}
