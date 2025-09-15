package lucide

import x "github.com/plainkit/html"

// ChevronsDown creates a Chevrons Down Lucide icon.
func ChevronsDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 6 5 5 5-5"))),
		x.Child(x.Path(x.D("m7 13 5 5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
