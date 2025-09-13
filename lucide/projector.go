package lucide

import x "github.com/bloxui/blox"

// Projector creates a Projector Lucide icon.
func Projector(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-projector", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 7 3 5"))),
		x.Child(x.Path(x.D("M9 6V3"))),
		x.Child(x.Path(x.D("m13 7 2-2"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("13"), x.R("3"))),
		x.Child(x.Path(x.D("M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17"))),
		x.Child(x.Path(x.D("M16 16h2"))),
	)
	return x.Svg(svgArgs...)
}
