package lucide

import x "github.com/plainkit/html"

// ChevronLeft creates a Chevron Left Lucide icon.
func ChevronLeft(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevron-left", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 18-6-6 6-6"))),
	)
	return x.Svg(svgArgs...)
}
