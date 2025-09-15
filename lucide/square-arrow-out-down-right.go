package lucide

import x "github.com/plainkit/blox"

// SquareArrowOutDownRight creates a Square Arrow Out Down Right Lucide icon.
func SquareArrowOutDownRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-square-arrow-out-down-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		x.Child(x.Path(x.D("m21 21-9-9"))),
		x.Child(x.Path(x.D("M21 15v6h-6"))),
	)
	return x.Svg(svgArgs...)
}
