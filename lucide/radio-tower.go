package lucide

import x "github.com/bloxui/blox"

// RadioTower creates a Radio Tower Lucide icon.
func RadioTower(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-radio-tower", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4.9 16.1C1 12.2 1 5.8 4.9 1.9"))),
		x.Child(x.Path(x.D("M7.8 4.7a6.14 6.14 0 0 0-.8 7.5"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("9"), x.R("2"))),
		x.Child(x.Path(x.D("M16.2 4.8c2 2 2.26 5.11.8 7.47"))),
		x.Child(x.Path(x.D("M19.1 1.9a9.96 9.96 0 0 1 0 14.1"))),
		x.Child(x.Path(x.D("M9.5 18h5"))),
		x.Child(x.Path(x.D("m8 22 4-11 4 11"))),
	)
	return x.Svg(svgArgs...)
}
