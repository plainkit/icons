package lucide

import x "github.com/plainkit/html"

// Refrigerator creates a Refrigerator Lucide icon.
func Refrigerator(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-refrigerator", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Z"))),
		x.Child(x.Path(x.D("M5 10h14"))),
		x.Child(x.Path(x.D("M15 7v6"))),
	)
	return x.Svg(svgArgs...)
}
