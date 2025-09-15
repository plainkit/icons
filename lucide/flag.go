package lucide

import x "github.com/plainkit/blox"

// Flag creates a Flag Lucide icon.
func Flag(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flag", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22V4a1 1 0 0 1 .4-.8A6 6 0 0 1 8 2c3 0 5 2 7.333 2q2 0 3.067-.8A1 1 0 0 1 20 4v10a1 1 0 0 1-.4.8A6 6 0 0 1 16 16c-3 0-5-2-8-2a6 6 0 0 0-4 1.528"))),
	)
	return x.Svg(svgArgs...)
}
