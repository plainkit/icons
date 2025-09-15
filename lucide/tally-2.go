package lucide

import x "github.com/plainkit/blox"

// Tally2 creates a Tally 2 Lucide icon.
func Tally2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tally-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 4v16"))),
		x.Child(x.Path(x.D("M9 4v16"))),
	)
	return x.Svg(svgArgs...)
}
