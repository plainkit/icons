package lucide

import x "github.com/bloxui/blox"

// MarsStroke creates a Mars Stroke Lucide icon.
func MarsStroke(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mars-stroke", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m14 6 4 4"))),
		x.Child(x.Path(x.D("M17 3h4v4"))),
		x.Child(x.Path(x.D("m21 3-7.75 7.75"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("15"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
