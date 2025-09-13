package lucide

import x "github.com/bloxui/blox"

// Copyleft creates a Copyleft Lucide icon.
func Copyleft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-copyleft", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M9.17 14.83a4 4 0 1 0 0-5.66"))),
	)
	return x.Svg(svgArgs...)
}
