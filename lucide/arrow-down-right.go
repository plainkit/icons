package lucide

import x "github.com/plainkit/blox"

// ArrowDownRight creates a Arrow Down Right Lucide icon.
func ArrowDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 7 10 10"))),
		x.Child(x.Path(x.D("M17 7v10H7"))),
	)
	return x.Svg(svgArgs...)
}
