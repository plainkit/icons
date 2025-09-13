package lucide

import x "github.com/bloxui/blox"

// Mars creates a Mars Lucide icon.
func Mars(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mars", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3h5v5"))),
		x.Child(x.Path(x.D("m21 3-6.75 6.75"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("14"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
