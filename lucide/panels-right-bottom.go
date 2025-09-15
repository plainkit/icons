package lucide

import x "github.com/plainkit/html"

// PanelsRightBottom creates a Panels Right Bottom Lucide icon.
func PanelsRightBottom(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panels-right-bottom", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M3 15h12"))),
		x.Child(x.Path(x.D("M15 3v18"))),
	)
	return x.Svg(svgArgs...)
}
