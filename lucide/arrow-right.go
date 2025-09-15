package lucide

import x "github.com/plainkit/blox"

// ArrowRight creates a Arrow Right Lucide icon.
func ArrowRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 12h14"))),
		x.Child(x.Path(x.D("m12 5 7 7-7 7"))),
	)
	return x.Svg(svgArgs...)
}
