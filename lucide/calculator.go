package lucide

import x "github.com/plainkit/html"

// Calculator creates a Calculator Lucide icon.
func Calculator(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-calculator", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("20"), x.X("4"), x.Y("2"), x.Rx("2"))),
		x.Child(x.Line(x.X1("8"), x.X2("16"), x.Y1("6"), x.Y2("6"))),
		x.Child(x.Line(x.X1("16"), x.X2("16"), x.Y1("14"), x.Y2("18"))),
		x.Child(x.Path(x.D("M16 10h.01"))),
		x.Child(x.Path(x.D("M12 10h.01"))),
		x.Child(x.Path(x.D("M8 10h.01"))),
		x.Child(x.Path(x.D("M12 14h.01"))),
		x.Child(x.Path(x.D("M8 14h.01"))),
		x.Child(x.Path(x.D("M12 18h.01"))),
		x.Child(x.Path(x.D("M8 18h.01"))),
	)
	return x.Svg(svgArgs...)
}
