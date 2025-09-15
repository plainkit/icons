package lucide

import x "github.com/plainkit/blox"

// ArrowUpRight creates a Arrow Up Right Lucide icon.
func ArrowUpRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 7h10v10"))),
		x.Child(x.Path(x.D("M7 17 17 7"))),
	)
	return x.Svg(svgArgs...)
}
