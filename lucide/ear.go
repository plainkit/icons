package lucide

import x "github.com/plainkit/blox"

// Ear creates a Ear Lucide icon.
func Ear(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ear", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 8.5a6.5 6.5 0 1 1 13 0c0 6-6 6-6 10a3.5 3.5 0 1 1-7 0"))),
		x.Child(x.Path(x.D("M15 8.5a2.5 2.5 0 0 0-5 0v1a2 2 0 1 1 0 4"))),
	)
	return x.Svg(svgArgs...)
}
