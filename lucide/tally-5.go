package lucide

import x "github.com/plainkit/html"

// Tally5 creates a Tally 5 Lucide icon.
func Tally5(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tally-5", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 4v16"))),
		x.Child(x.Path(x.D("M9 4v16"))),
		x.Child(x.Path(x.D("M14 4v16"))),
		x.Child(x.Path(x.D("M19 4v16"))),
		x.Child(x.Path(x.D("M22 6 2 18"))),
	)
	return x.Svg(svgArgs...)
}
