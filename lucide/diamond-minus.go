package lucide

import x "github.com/plainkit/html"

// DiamondMinus creates a Diamond Minus Lucide icon.
func DiamondMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-diamond-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0z"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
