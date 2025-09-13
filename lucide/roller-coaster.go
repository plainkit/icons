package lucide

import x "github.com/bloxui/blox"

// RollerCoaster creates a Roller Coaster Lucide icon.
func RollerCoaster(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-roller-coaster", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 19V5"))),
		x.Child(x.Path(x.D("M10 19V6.8"))),
		x.Child(x.Path(x.D("M14 19v-7.8"))),
		x.Child(x.Path(x.D("M18 5v4"))),
		x.Child(x.Path(x.D("M18 19v-6"))),
		x.Child(x.Path(x.D("M22 19V9"))),
		x.Child(x.Path(x.D("M2 19V9a4 4 0 0 1 4-4c2 0 4 1.33 6 4s4 4 6 4a4 4 0 1 0-3-6.65"))),
	)
	return x.Svg(svgArgs...)
}
