package lucide

import x "github.com/plainkit/html"

// ListFilter creates a List Filter Lucide icon.
func ListFilter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-filter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 5h20"))),
		x.Child(x.Path(x.D("M6 12h12"))),
		x.Child(x.Path(x.D("M9 19h6"))),
	)
	return x.Svg(svgArgs...)
}
