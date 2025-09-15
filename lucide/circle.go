package lucide

import x "github.com/plainkit/blox"

// Circle creates a Circle Lucide icon.
func Circle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
	)
	return x.Svg(svgArgs...)
}
