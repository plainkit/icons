package lucide

import x "github.com/bloxui/blox"

// Truck creates a Truck Lucide icon.
func Truck(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-truck", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2"))),
		x.Child(x.Path(x.D("M15 18H9"))),
		x.Child(x.Path(x.D("M19 18h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.624l-3.48-4.35A1 1 0 0 0 17.52 8H14"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("18"), x.R("2"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("18"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
