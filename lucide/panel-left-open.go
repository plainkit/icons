package lucide

import x "github.com/plainkit/html"

// PanelLeftOpen creates a Panel Left Open Lucide icon.
func PanelLeftOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-left-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Path(x.D("m14 9 3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
