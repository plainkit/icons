package lucide

import x "github.com/bloxui/blox"

// Proportions creates a Proportions Lucide icon.
func Proportions(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-proportions", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("20"), x.RectHeight("16"), x.X("2"), x.Y("4"), x.Rx("2"))),
		x.Child(x.Path(x.D("M12 9v11"))),
		x.Child(x.Path(x.D("M2 9h13a2 2 0 0 1 2 2v9"))),
	)
	return x.Svg(svgArgs...)
}
