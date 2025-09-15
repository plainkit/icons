package lucide

import x "github.com/plainkit/blox"

// Pocket creates a Pocket Lucide icon.
func Pocket(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pocket", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 3a2 2 0 0 1 2 2v6a1 1 0 0 1-20 0V5a2 2 0 0 1 2-2z"))),
		x.Child(x.Path(x.D("m8 10 4 4 4-4"))),
	)
	return x.Svg(svgArgs...)
}
