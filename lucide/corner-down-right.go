package lucide

import x "github.com/plainkit/blox"

// CornerDownRight creates a Corner Down Right Lucide icon.
func CornerDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-corner-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15 10 5 5-5 5"))),
		x.Child(x.Path(x.D("M4 4v7a4 4 0 0 0 4 4h12"))),
	)
	return x.Svg(svgArgs...)
}
