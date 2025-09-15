package lucide

import x "github.com/plainkit/blox"

// ArrowDown creates a Arrow Down Lucide icon.
func ArrowDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5v14"))),
		x.Child(x.Path(x.D("m19 12-7 7-7-7"))),
	)
	return x.Svg(svgArgs...)
}
