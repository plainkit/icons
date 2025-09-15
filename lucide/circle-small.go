package lucide

import x "github.com/plainkit/blox"

// CircleSmall creates a Circle Small Lucide icon.
func CircleSmall(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-small", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
