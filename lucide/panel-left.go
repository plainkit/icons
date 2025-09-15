package lucide

import x "github.com/plainkit/html"

// PanelLeft creates a Panel Left Lucide icon.
func PanelLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panel-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 3v18"))),
	)
	return x.Svg(svgArgs...)
}
