package lucide

import x "github.com/plainkit/html"

// ListMinus creates a List Minus Lucide icon.
func ListMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 5H3"))),
		x.Child(x.Path(x.D("M11 12H3"))),
		x.Child(x.Path(x.D("M16 19H3"))),
		x.Child(x.Path(x.D("M21 12h-6"))),
	)
	return x.Svg(svgArgs...)
}
