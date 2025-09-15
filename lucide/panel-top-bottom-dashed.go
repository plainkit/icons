package lucide

import x "github.com/plainkit/html"

// PanelTopBottomDashed creates a Panel Top Bottom Dashed Lucide icon.
func PanelTopBottomDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-top-bottom-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 15h1"))),
		x.Child(x.Path(x.D("M14 9h1"))),
		x.Child(x.Path(x.D("M19 15h2"))),
		x.Child(x.Path(x.D("M19 9h2"))),
		x.Child(x.Path(x.D("M3 15h2"))),
		x.Child(x.Path(x.D("M3 9h2"))),
		x.Child(x.Path(x.D("M9 15h1"))),
		x.Child(x.Path(x.D("M9 9h1"))),
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
