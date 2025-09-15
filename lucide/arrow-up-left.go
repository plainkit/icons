package lucide

import x "github.com/plainkit/blox"

// ArrowUpLeft creates a Arrow Up Left Lucide icon.
func ArrowUpLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 17V7h10"))),
		x.Child(x.Path(x.D("M17 17 7 7"))),
	)
	return x.Svg(svgArgs...)
}
