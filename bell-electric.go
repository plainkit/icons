package lucide

import x "github.com/bloxui/blox"

// BellElectric creates a Bell Electric Lucide icon.
func BellElectric(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bell-electric", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.518 17.347A7 7 0 0 1 14 19"))),
		x.Child(x.Path(x.D("M18.8 4A11 11 0 0 1 20 9"))),
		x.Child(x.Path(x.D("M9 9h.01"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("16"), x.R("2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("9"), x.R("7"))),
		x.Child(x.Rect(x.X("4"), x.Y("16"), x.RectWidth("10"), x.RectHeight("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
