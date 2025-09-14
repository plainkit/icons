package lucide

import x "github.com/bloxui/blox"

// Dot creates a Dot Lucide icon.
func Dot(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-dot", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12.1"), x.Cy("12.1"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
