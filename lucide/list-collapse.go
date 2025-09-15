package lucide

import x "github.com/plainkit/html"

// ListCollapse creates a List Collapse Lucide icon.
func ListCollapse(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-collapse", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 5h11"))),
		x.Child(x.Path(x.D("M10 12h11"))),
		x.Child(x.Path(x.D("M10 19h11"))),
		x.Child(x.Path(x.D("m3 10 3-3-3-3"))),
		x.Child(x.Path(x.D("m3 20 3-3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
