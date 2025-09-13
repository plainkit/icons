package lucide

import x "github.com/bloxui/blox"

// Rows4 creates a Rows 4 Lucide icon.
func Rows4(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-rows-4", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("18"), x.RectHeight("18"), x.X("3"), x.Y("3"), x.Rx("2"))),
		x.Child(x.Path(x.D("M21 7.5H3"))),
		x.Child(x.Path(x.D("M21 12H3"))),
		x.Child(x.Path(x.D("M21 16.5H3"))),
	)
	return x.Svg(svgArgs...)
}
