package lucide

import x "github.com/plainkit/blox"

// Beaker creates a Beaker Lucide icon.
func Beaker(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-beaker", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4.5 3h15"))),
		x.Child(x.Path(x.D("M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3"))),
		x.Child(x.Path(x.D("M6 14h12"))),
	)
	return x.Svg(svgArgs...)
}
