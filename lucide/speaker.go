package lucide

import x "github.com/bloxui/blox"

// Speaker creates a Speaker Lucide icon.
func Speaker(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-speaker", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 6h.01"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("14"), x.R("4"))),
		x.Child(x.Path(x.D("M12 14h.01"))),
	)
	return x.Svg(svgArgs...)
}
