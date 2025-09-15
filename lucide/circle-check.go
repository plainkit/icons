package lucide

import x "github.com/plainkit/blox"

// CircleCheck creates a Circle Check Lucide icon.
func CircleCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("m9 12 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
