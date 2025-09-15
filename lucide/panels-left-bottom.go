package lucide

import x "github.com/plainkit/blox"

// PanelsLeftBottom creates a Panels Left Bottom Lucide icon.
func PanelsLeftBottom(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-panels-left-bottom", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M9 3v18"))),
		x.Child(x.Path(x.D("M9 15h12"))),
	)
	return x.Svg(svgArgs...)
}
