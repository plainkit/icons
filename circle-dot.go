package lucide

import x "github.com/bloxui/blox"

// CircleDot creates a Circle Dot Lucide icon.
func CircleDot(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-circle-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}