package lucide

import x "github.com/bloxui/blox"

// Blinds creates a Blinds Lucide icon.
func Blinds(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-blinds", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 3h18"))),
		x.Child(x.Path(x.D("M20 7H8"))),
		x.Child(x.Path(x.D("M20 11H8"))),
		x.Child(x.Path(x.D("M10 19h10"))),
		x.Child(x.Path(x.D("M8 15h12"))),
		x.Child(x.Path(x.D("M4 3v14"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("19"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
