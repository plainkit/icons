package lucide

import x "github.com/plainkit/blox"

// IterationCcw creates a Iteration Ccw Lucide icon.
func IterationCcw(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-iteration-ccw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 14 4 4-4 4"))),
		x.Child(x.Path(x.D("M20 10a8 8 0 1 0-8 8h8"))),
	)
	return x.Svg(svgArgs...)
}
