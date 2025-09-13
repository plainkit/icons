package lucide

import x "github.com/bloxui/blox"

// CableCar creates a Cable Car Lucide icon.
func CableCar(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cable-car", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 3h.01"))),
		x.Child(x.Path(x.D("M14 2h.01"))),
		x.Child(x.Path(x.D("m2 9 20-5"))),
		x.Child(x.Path(x.D("M12 12V6.5"))),
		x.Child(x.Rect(x.X("4"), x.Y("12"), x.RectWidth("16"), x.RectHeight("10"), x.Rx("3"))),
		x.Child(x.Path(x.D("M9 12v5"))),
		x.Child(x.Path(x.D("M15 12v5"))),
		x.Child(x.Path(x.D("M4 17h16"))),
	)
	return x.Svg(svgArgs...)
}
