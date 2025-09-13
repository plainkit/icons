package lucide

import x "github.com/bloxui/blox"

// IterationCw creates a Iteration Cw Lucide icon.
func IterationCw(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-iteration-cw", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 10a8 8 0 1 1 8 8H4"))),
		x.Child(x.Path(x.D("m8 22-4-4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
