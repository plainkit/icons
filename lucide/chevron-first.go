package lucide

import x "github.com/plainkit/html"

// ChevronFirst creates a Chevron First Lucide icon.
func ChevronFirst(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-chevron-first", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m17 18-6-6 6-6"))),
		x.Child(x.Path(x.D("M7 6v12"))),
	)
	return x.Svg(svgArgs...)
}
