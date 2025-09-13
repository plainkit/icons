package lucide

import x "github.com/bloxui/blox"

// Instagram creates an Instagram Lucide icon.
func Instagram(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-instagram", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.X("2"), x.Y("2"), x.RectWidth("20"), x.RectHeight("20"), x.Rx("5"), x.Ry("5"))),
		x.Child(x.Path(x.D("M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"))),
		x.Child(x.Line(x.X1("17.5"), x.Y1("6.5"), x.X2("17.51"), x.Y2("6.5"))),
	)
	return x.Svg(svgArgs...)
}
