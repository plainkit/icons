package lucide

import x "github.com/plainkit/html"

// Baseline creates a Baseline Lucide icon.
func Baseline(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-baseline", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 20h16"))),
		x.Child(x.Path(x.D("m6 16 6-12 6 12"))),
		x.Child(x.Path(x.D("M8 12h8"))),
	)
	return x.Svg(svgArgs...)
}
