package lucide

import x "github.com/bloxui/blox"

// PcCase creates a Pc Case Lucide icon.
func PcCase(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pc-case", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("14"), x.RectHeight("20"), x.X("5"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M15 14h.01"))),
		x.Child(x.Path(x.D("M9 6h6"))),
		x.Child(x.Path(x.D("M9 10h6"))),
	)
	return x.Svg(svgArgs...)
}
