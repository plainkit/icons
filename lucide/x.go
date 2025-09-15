package lucide

import x "github.com/plainkit/blox"

// X creates a X Lucide icon.
func X(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 6 6 18"))),
		x.Child(x.Path(x.D("m6 6 12 12"))),
	)
	return x.Svg(svgArgs...)
}
