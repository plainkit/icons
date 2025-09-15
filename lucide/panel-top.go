package lucide

import x "github.com/plainkit/html"

// PanelTop creates a Panel Top Lucide icon.
func PanelTop(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-top", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
	)
	return x.Svg(svgArgs...)
}
