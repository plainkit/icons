package lucide

import x "github.com/bloxui/blox"

// Stethoscope creates a Stethoscope Lucide icon.
func Stethoscope(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-stethoscope", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 2v2"))),
		x.Child(x.Path(x.D("M5 2v2"))),
		x.Child(x.Path(x.D("M5 3H4a2 2 0 0 0-2 2v4a6 6 0 0 0 12 0V5a2 2 0 0 0-2-2h-1"))),
		x.Child(x.Path(x.D("M8 15a6 6 0 0 0 12 0v-3"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("10"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
