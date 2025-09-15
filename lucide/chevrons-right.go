package lucide

import x "github.com/plainkit/html"

// ChevronsRight creates a Chevrons Right Lucide icon.
func ChevronsRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevrons-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m6 17 5-5-5-5"))),
		x.Child(x.Path(x.D("m13 17 5-5-5-5"))),
	)
	return x.Svg(svgArgs...)
}
