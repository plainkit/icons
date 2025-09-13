package lucide

import x "github.com/bloxui/blox"

// TruckElectric creates a Truck Electric Lucide icon.
func TruckElectric(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-truck-electric", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 19V7a2 2 0 0 0-2-2H9"))),
		x.Child(x.Path(x.D("M15 19H9"))),
		x.Child(x.Path(x.D("M19 19h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.62L18.3 9.38a1 1 0 0 0-.78-.38H14"))),
		x.Child(x.Path(x.D("M2 13v5a1 1 0 0 0 1 1h2"))),
		x.Child(x.Path(x.D("M4 3 2.15 5.15a.495.495 0 0 0 .35.86h2.15a.47.47 0 0 1 .35.86L3 9.02"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("19"), x.R("2"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
