package lucide

import x "github.com/bloxui/blox"

// CircleArrowDown creates a Circle Arrow Down Lucide icon.
func CircleArrowDown(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M12 8v8"))),
		x.Child(x.Path(x.D("m8 12 4 4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
