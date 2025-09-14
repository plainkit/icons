package lucide

import x "github.com/bloxui/blox"

// ListIndentDecrease creates a List Indent Decrease Lucide icon.
func ListIndentDecrease(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-list-indent-decrease", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 5H11"))),
		x.Child(x.Path(x.D("M21 12H11"))),
		x.Child(x.Path(x.D("M21 19H11"))),
		x.Child(x.Path(x.D("m7 8-4 4 4 4"))),
	)
	return x.Svg(svgArgs...)
}
