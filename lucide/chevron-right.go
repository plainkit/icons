package lucide

import x "github.com/plainkit/html"

// ChevronRight creates a Chevron Right Lucide icon.
func ChevronRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevron-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 18 6-6-6-6"))),
	)
	return x.Svg(svgArgs...)
}
