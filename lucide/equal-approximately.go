package lucide

import x "github.com/plainkit/html"

// EqualApproximately creates a Equal Approximately Lucide icon.
func EqualApproximately(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-equal-approximately", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 15a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"))),
		x.Child(x.Path(x.D("M5 9a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"))),
	)
	return x.Svg(svgArgs...)
}
