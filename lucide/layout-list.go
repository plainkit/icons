package lucide

import x "github.com/bloxui/blox"

// LayoutList creates a Layout List Lucide icon.
func LayoutList(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-layout-list", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("7"), x.X("3"), x.Y("14"), x.Rx("1"))),
		x.Child(x.Path(x.D("M14 4h7"))),
		x.Child(x.Path(x.D("M14 9h7"))),
		x.Child(x.Path(x.D("M14 15h7"))),
		x.Child(x.Path(x.D("M14 20h7"))),
	)
	return x.Svg(svgArgs...)
}
