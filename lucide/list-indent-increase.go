package lucide

import x "github.com/plainkit/blox"

// ListIndentIncrease creates a List Indent Increase Lucide icon.
func ListIndentIncrease(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-indent-increase", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H11"))),
		x.Child(x.Path(x.D("M21 12H11"))),
		x.Child(x.Path(x.D("M21 19H11"))),
		x.Child(x.Path(x.D("m3 8 4 4-4 4"))),
	)
	return x.Svg(svgArgs...)
}
