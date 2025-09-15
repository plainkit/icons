package lucide

import x "github.com/plainkit/html"

// PanelTopOpen creates a Panel Top Open Lucide icon.
func PanelTopOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-top-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 9h18"))),
		x.Child(x.Path(x.D("m15 14-3 3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
