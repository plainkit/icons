package lucide

import x "github.com/bloxui/blox"

// CircleArrowLeft creates a Circle Arrow Left Lucide icon.
func CircleArrowLeft(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m12 8-4 4 4 4"))),
		x.Child(x.Path(x.D("M16 12H8"))),
	)
	return x.Svg(svgArgs...)
}
