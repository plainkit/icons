package lucide

import x "github.com/bloxui/blox"

// ShipWheel creates a Ship Wheel Lucide icon.
func ShipWheel(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ship-wheel", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("8"))),
		x.Child(x.Path(x.D("M12 2v7.5"))),
		x.Child(x.Path(x.D("m19 5-5.23 5.23"))),
		x.Child(x.Path(x.D("M22 12h-7.5"))),
		x.Child(x.Path(x.D("m19 19-5.23-5.23"))),
		x.Child(x.Path(x.D("M12 14.5V22"))),
		x.Child(x.Path(x.D("M10.23 13.77 5 19"))),
		x.Child(x.Path(x.D("M9.5 12H2"))),
		x.Child(x.Path(x.D("M10.23 10.23 5 5"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("2.5"))),
	)
	return x.Svg(svgArgs...)
}
